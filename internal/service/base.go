package service

import (
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"slices"
	"strings"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/internal/service/user"
	"xiaoniuds.com/cid/pkg/errs"
	msgSender "xiaoniuds.com/cid/pkg/msg_sender"
	"xiaoniuds.com/cid/vars"
)

type Base struct {
	C         *config.Config
	DbConnect *data.Data
}

// CidSendNotify Cid 发送通知消息
func (s *Base) CidSendNotify(mainUserId int64, msgFmt string, msgParams []interface{}, notifyType string, dataOwnerUserId int64) (err *errs.MyErr) {
	builder := func(query *gorm.DB) *gorm.DB {
		return query.Where("main_user_id = ? and notify_type = ?", mainUserId, notifyType)
	}
	settings, err := common.NewNotifySettingModel("", s.DbConnect).GetNotifySettingListByBuilder(builder, []string{})
	if err != nil {
		return
	}
	for _, setting := range settings {
		// 设置的权限级别的需要判断推送的通知所属
		if !s.infoIsNotifyByLevel(dataOwnerUserId, setting) {
			continue
		}
		// 检查类型通知检查配置
		if !s.infoIsNotifyByType(notifyType, setting, msgFmt) {
			continue
		}
		//try {
		webhooks := []string{setting.Webhook}
		// 短信发送的时候，可指定多个手机号发送
		if setting.NotifyMethod == "sms" {
			webhooks = strings.Split(setting.Webhook, ",")
		}
		sender := msgSender.NewMessageSender(setting.NotifyMethod, webhooks)
		// 判断是否为钉钉消息，如果是可以换成markdown格式，并且附带商品预览图
		var (
			markdown *msgSender.MarkdownMsg
			message  *msgSender.NotifyMsg
		)
		markdown, err = s.cidFormatNotify(mainUserId, msgFmt, msgParams, setting.NotifyMethod)
		if err != nil {
			// log
			vars.SysLog.Warnf("cidFormatNotify data:%v,msg:%s", setting, err.Error())
			continue
		}
		if markdown == nil {
			message = sender.TextMessage(&msgSender.TextMsg{
				MsgFmt: msgFmt,
				Params: msgParams,
			})
		} else {
			message = sender.MarkdownMessage(markdown)
		}
		err = sender.Send(message)
		//	} catch (\Throwable $throwable) {
		//	LogService::error(['notify_data' => $notifyDatum, 'msg' => $throwable->getMessage(),'notifyMethod'=>$notifyDatum['notify_method'],'webhooks'=>$webhooks,'sendMsg'=>$msg,'msgParams'=>$msgParams,'markdownArr'=>$markdownArr], 'cid_send_notify');
		//}
	}

	return
}

// 检查本次通知的对象是否符合配置的等级
func (s *Base) infoIsNotifyByLevel(dataOwnerUserId int64, notifyDatum *common.NotifySetting) bool {
	if dataOwnerUserId == 0 || notifyDatum.NotifyLevel == vars.NotifyLevelCompany {
		// 设置的企业级别默认推送
		return true
	}

	return s.promotionLimitNotify(dataOwnerUserId, notifyDatum)
}

func (s *Base) promotionLimitNotify(dataOwnerUserId int64, notifyDatum *common.NotifySetting) bool {
	if notifyDatum.NotifyLevel == vars.NotifyLevelDepartment {
		users, err := (&user.Service{C: s.C, DbConnect: s.DbConnect}).
			GetMyAuthorizedUsers("default", "data_range", notifyDatum.OwnerUserId, 2, -1)
		if err != nil {
			return false
		}
		ownerIds := make([]int64, 0)
		for _, u := range users {
			ownerIds = append(ownerIds, u.UserId)
		}
		return slices.Contains(ownerIds, dataOwnerUserId)
	} else if notifyDatum.NotifyLevel == vars.NotifyLevelPerson {
		// 判断商品所属人是否是通知添加人
		return dataOwnerUserId == notifyDatum.OwnerUserId
	} else {
		return false
	}
}

// 类型拆分了子配置后，检查是否符合配置
func (s *Base) infoIsNotifyByType(notifyType string, notifyDatum *common.NotifySetting, templateKey string) bool {
	if notifyType == vars.NotifyTypeGoodsPromotionLimit {
		v, ok := vars.NotifyTypeGoodsPromotionLimitValue[templateKey]
		if !ok {
			return false
		}
		// 兼容历史库里的数据，默认都发
		if notifyDatum.ExtendedValue == "" {
			return true
		}
		var values []int
		err := json.Unmarshal([]byte(notifyDatum.ExtendedValue), &values)
		if err != nil {
			return false
		}

		return slices.Contains(values, v)
	}

	return true
}

/**
 * 构建markdown文本
 * @param $mainUserId
 * @param $msg
 * @param $msgParams
 * @param $notifyMethod
 * @return array
 */
func (s *Base) cidFormatNotify(mainUserId int64, msgFmt string, msgParams []interface{}, notifyMethod string) (markdown *msgSender.MarkdownMsg, err *errs.MyErr) {
	if mainUserId == 0 {
		return
	}
	if notifyMethod != "dingding_robot" {
		return
	}
	//根据模板内容，反向获取模板类型
	msgTemplateKey := "" //array_search(msgFmt, msgTemplate)
	if !slices.Contains([]string{
		"authSuccessNotifyTemplate",
		"authFailNotifyTemplate",
		"promotionLimitNotifyTemplate",
		"promotionNotifyTemplate",
	}, msgTemplateKey) {
		return
	}

	//text := msgSender.TextMsg{
	//	MsgFmt: msgFmt,
	//	Params: msgParams,
	//}
	msgText := fmt.Sprintf(msgFmt, msgParams...)
	textArr := strings.Split(msgText, "\n")
	title := textArr[0]

	if !slices.Contains([]string{
		"authFailNotifyTemplate",
		"promotionLimitNotifyTemplate",
	}, msgTemplateKey) {
		textArr[0] = fmt.Sprintf("<font color='red'>%s</font>", title)
	} else {
		textArr[0] = fmt.Sprintf("<font color='green'>%s</font>", title)
	}
	//
	////商品ID
	//$goodsId = intval($msgParams[3]);
	//$where = [
	//'main_user_id' => $mainUserId,
	//'goods_id' => $goodsId,
	//];
	//$goodsDetail = \App\Services\Cid\Pinduoduo\GoodsService::getGoodsOneByParam($where, ['goods_thumbnail_url','owner_user_id']);
	//if(empty($goodsDetail) || empty($goodsDetail['goods_thumbnail_url'])){
	//return [];
	//}
	//if($notifyMethod == 'dingding_robot'){
	//$textArr[] = "![预览图]({$goodsDetail['goods_thumbnail_url']})";
	//$markdown = implode("\n\n", $textArr);
	//}else{
	//return [];
	//}
	//$result = [
	//'title' => $title,
	//'markdown' => $markdown,
	//'goods_thumbnail_url' => $goodsDetail['goods_thumbnail_url']
	//];
	//LogService::info(['result' => $result, 'main_user_id' => $mainUserId], 'cidFormatNotify');
	//return $result;
	return
}
