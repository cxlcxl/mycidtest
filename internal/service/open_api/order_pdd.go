package open_api

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type PDD struct {
	ShopPayTimeField string
	VerifyTimeField  string
}

func (p *PDD) GetOrderList(params statement.OrderList, connects *data.Data) (orders []*OrderItem, total int64, err *errs.MyErr) {
	//params.OpenApiLoginData.MainUserId
	sort := p.getSortField(params.TimeType)
	// 限制只能查看企业 main_user_id
	advWheres := []string{
		fmt.Sprintf("ad_act.main_user_id = %d", params.OpenApiLoginData.MainUserId),
	}
	//是否直推
	//if (!empty($conditions['isDirect'])) {
	//$advWheres[] = "ad_act.is_direct = {$conditions['isDirect']}";
	//}
	//
	////是否隐藏订单（=是否跨店，此处跨店由小牛定义）
	//if (isset($conditions['isHidden'])) {
	//$isHidden = $conditions['isHidden'];
	//$advWheres[] = "ad_act.is_hide = {$isHidden}";
	//} else {
	//// $advWheres[] = "ad_act.no_callback_reason != '跨店订单不回传'";
	//}
	//if (!empty($conditions['type']) && $conditions['type'] > 0) {
	//$advWheres[] = "ad_act.type = {$conditions['type']}";
	//}
	//if (!empty($conditions['pid'])) {
	//$advWheres[] = "ad_act.p_id = '{$conditions['pid']}'";
	//}
	//if (!empty($conditions['mainUserIds'])) {
	//$advWheres[] = "ad_act.main_user_id in ({$conditions['mainUserIds']})";
	//}

	//// 指定查看用户
	//if (!empty($conditions['owner_user_ids'])) {
	//if (is_array($conditions['owner_user_ids'])) {
	//$conditions['owner_user_ids'] = implode(',', $conditions['owner_user_ids']);
	//}
	//$advWheres[] = "ad_act.owner_user_id in ({$conditions['owner_user_ids']})";
	//}
	//
	////订单状态
	//if ($conditions['order_status'] != "") {
	//$advWheres[] = "ad_act.order_status in ({$conditions ['order_status']})";
	//}
	//if ($conditions['id'] > 0) {
	//$id = (int)$conditions['id'];
	//$advWheres[] = "ad_act.id > {$id}";
	//}
	//
	////订单号
	//if (!empty($conditions['order_sn'])) {
	//$advWheres[] = "ad_act.order_sn ='{$conditions ['order_sn']}'";
	//}
	//
	////广告ID
	//if (!empty($conditions['ad_id'])) {
	//$conditions['ad_id'] = intval($conditions['ad_id']);
	//$advWheres[] = "ad_act.ad_id = {$conditions ['ad_id']} ";
	//}
	//
	//// 创量跟踪类型
	//if (!empty($conditions['cl_trace_type'])) {
	//$advWheres[] = "ad_act.cl_trace_type ={$conditions ['cl_trace_type']}";
	//}

	////商品
	//if (!empty($conditions['goods_keyword'])) {
	//if (is_numeric($conditions ['goods_keyword'])) {
	//$advWheres[] = "ad_act.goods_id ={$conditions ['goods_keyword']}";
	//} else {
	//$advWheres[] = "ad_act.goods_name like '%{$conditions ['goods_keyword']}%'";
	//}
	//}
	////推广位
	//if (!empty($conditions['promotion_keyword'])) {
	//$promotionKeyword = str_replace('_', '', $conditions['promotion_keyword']);
	//if (is_numeric($promotionKeyword)) {
	//$advWheres[] = "ad_act.p_id ='{$conditions ['promotion_keyword']}'";
	//} else {
	//$advWheres[] = "ad_act.p_name like '%{$conditions ['promotion_keyword']}%'";
	//}
	//}
	////创建时间
	//if (!empty($conditions['create_time'])) {
	//$tTime = explode(',', $conditions['create_time']);
	//$startDate = $tTime[0] ?? date('Y-m-d  H:i:s');
	//$endDate = $tTime[1] ?? date('Y-m-d  H:i:s');
	//$advWheres[] = "(ad_act.create_time >= '{$startDate}' and ad_act.create_time <='{$endDate}')";
	//}
	////更新时间
	//if (!empty($conditions['update_time'])) {
	//$tTime = explode(',', $conditions['update_time']);
	//$startDate = $tTime[0] ?? date('Y-m-d  H:i:s');
	//$endDate = $tTime[1] ?? date('Y-m-d  H:i:s');
	//$advWheres[] = "(ad_act.update_time >= '{$startDate}' and ad_act.update_time <='{$endDate}')";
	//}
	////支付时间
	//if (!empty($conditions['order_pay_time'])) {
	//$tTime = explode(',', $conditions['order_pay_time']);
	//$startDate = strtotime($tTime[0] ?? date('Y-m-d  H:i:s'));
	//$endDate = strtotime($tTime[1] ?? date('Y-m-d H:i:s'));
	//$advWheres[] = "(ad_act.order_pay_time >= '{$startDate}' and ad_act.order_pay_time <='{$endDate}')";
	//}
	////审核时间
	//if (!empty($conditions['order_verify_time'])) {
	//$tTime = explode(',', $conditions['order_verify_time']);
	//$startDate = strtotime($tTime[0] ?? date('Y-m-d  H:i:s'));
	//$endDate = strtotime($tTime[1] ?? date('Y-m-d H:i:s'));
	//$advWheres[] = "(ad_act.order_verify_time >= '{$startDate}' and ad_act.order_verify_time <='{$endDate}')";
	//}
	//// 夸品订单查询
	//if (isset($conditions['is_diff_goods']) && is_numeric($conditions['is_diff_goods'])) {
	//$advWheres[] = "ad_act.is_diff_goods ={$conditions['is_diff_goods']}";
	//}
	////时间限制 1天
	//if (isset($endDate) && isset($startDate) && (($endDate - $startDate) / 86400) > 1) {
	//return ['list' => []];
	//}
	////多多进宝账号ID
	//if (!empty($conditions['account_user_id'])) {
	//$aIds = $conditions['account_user_id'];
	//$advWheres[] = " ad_act.account_user_id in ({$aIds}) ";
	//}
	//// 回传状态
	///*
	// * 0-未回传
	// * 1-已回传
	// * delay_callback-延迟回传中 (对应查询未回传原因为"延迟回传"的订单)
	// */
	//if (isset($conditions['is_callback']) && is_numeric($conditions['is_callback'])) {
	//$isCallback = $conditions['is_callback'] > 0 ? $conditions['is_callback'] : 0;
	//$advWheres[] = " ad_act.is_callback ={$isCallback} ";
	//}
	//$noCallbackReason = [];
	//if (isset($conditions['is_callback']) && $conditions['is_callback'] == 'delay_callback') {
	//$noCallbackReason[] = '延迟回传';
	//}
	//// 未回传原因
	//if (!empty($conditions['no_callback_reason'])) {
	//$noCallbackReason[] = $conditions['no_callback_reason'];
	//}
	//if (!empty($noCallbackReason)) {
	//// 订单内商品数量大于N不回传 特殊处理
	//if (in_array('订单内商品数量大于N不回传', $noCallbackReason)) {
	//$advWheres[] = " ad_act.no_callback_reason like '订单内商品数量大于%不回传' ";
	//unset($noCallbackReason[array_search('订单内商品数量大于N不回传', $noCallbackReason)]);
	//}
	//// 如果还有数据未回传原因，则按照原有逻辑处理
	//if (!empty($noCallbackReason)) {
	//$noCallbackReason = implode("','", $noCallbackReason);
	//$advWheres[] = " ad_act.no_callback_reason in ('{$noCallbackReason}') ";
	//}
	//}
	//// 回传类型
	//if (!empty($conditions['callback_type']) || is_numeric($conditions['callback_type'])) {
	//$callbackType = $conditions['callback_type'];
	//$advWheres[] = " ad_act.callback_type ={$callbackType} ";
	//}
	////归因状态
	//if (!empty($conditions['is_click']) || is_numeric($conditions['is_click'])) {
	//$isClick = $conditions['is_click'] > 0 ? 1 : 0;
	//if ($isClick == 0) {
	//$advWheres[] = " ad_act.click_id = 0 ";
	//} else {
	//$advWheres[] = " ad_act.click_id > 0 ";
	//}
	//}
	//// 媒体类型
	//if (!empty($conditions['media_type']) || is_numeric($conditions['media_type'])) {
	//$mediaType = empty($conditions['media_type']) ? 0 : $conditions['media_type'];
	//$advWheres[] = " ad_act.media_type in ({$mediaType})";
	//}
	////广告账号ID
	//if (!empty($conditions['advertiser_id'])) {
	//$ttaIds = $conditions['advertiser_id'];
	//$advWheres[] = " ad_act.advertiser_id in ({$ttaIds}) ";
	//}
	//// 店铺
	//if (!empty($conditions['mall_ids'])) {
	//$advWheres[] = " ad_act.mall_id in ({$conditions['mall_ids']})";
	//}
	//
	baseFields := []string{
		"ad_act.id", "ad_act.order_sn", "ad_act.p_id", "ad_act.p_name", "ad_act.is_hide", "ad_act.goods_id",
		"ad_act.goods_name", "ad_act.goods_quantity", "ad_act.promotion_rate", "ad_act.goods_thumbnail_url", "ad_act.order_amount",
		"ad_act.mall_id", "ad_act.mall_name", "ad_act.order_status", "ad_act.order_status_desc", "ad_act.fail_reason",
		"ad_act.type", "ad_act.click_id", "ad_act.ad_id", "ad_act.ad_site_id", "ad_act.media_type", "ad_act.account_user_id",
		"ad_act.is_callback", "ad_act.no_callback_reason", "ad_act.cl_trace_type", "ad_act.is_direct", "ad_act.callback_type",
		"ad_act.owner_user_id", "ad_act.promotion_type", "ad_act.is_diff_goods", "ad_act.custom_parameters",
		"ad_act.csite_type", "ad_act.callback_event", "ad_act.main_user_id", "ad_act.create_time", "ad_act.update_time",
	}
	groupByFields := append(
		baseFields,
		"ad_act.promotion_amount",
		"ad_act.goods_price",
		"ad_act.advertiser_id,med.advertiser_nick",
		"ad_act.id",
	)
	// 字段异常问题
	selectFields := append(
		baseFields,
		"COALESCE(ad_act.promotion_amount, 0) as promotion_amount",
		"COALESCE(ad_act.goods_price, 0) as goods_price",
		"CASE WHEN COALESCE(max(ad_act.order_pay_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_pay_time)))  END as order_pay_time",
		"CASE WHEN COALESCE(max(ad_act.order_create_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_create_time)))  END as order_create_time",
		"CASE WHEN COALESCE(max(ad_act.order_verify_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_verify_time)))  END as order_verify_time",
		"CASE WHEN COALESCE(max(ad_act.order_receive_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_receive_time)))  END as order_receive_time",
		"CASE WHEN COALESCE(max(ad_act.callback_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.callback_time)))  END as callback_time",
		//order_modify_at字段下线，使用退款回传时间充当退款时间
		"CASE WHEN COALESCE(max(ad_act.order_status),0) != 4 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.refund_callback_time)))  END as order_refund_time",
		"ad_act.advertiser_id,med.advertiser_nick",
		"MAX(t_admin_user.user_name) as user_name",
		"MAX(t_admin_user.user_full_name) as user_full_name",
	)
	leftJoins := []string{
		"LEFT JOIN chuangliang_doris_common.media_account AS med ON ad_act.advertiser_id = med.advertiser_id and ad_act.main_user_id = med.main_user_id and med.is_delete=0 and med.advertiser_id>0",
		"LEFT JOIN chuangliang_doris_common.admin_user t_admin_user ON ad_act.main_user_id = t_admin_user.main_user_id and t_admin_user.parent_id = 0",
	}

	selectFieldSQL := strings.Join(selectFields, ",")
	groupByFieldSQL := strings.Join(groupByFields, ",")
	advWhereSQL := strings.Join(advWheres, " AND ")
	leftJoinSQL := strings.Join(leftJoins, " ")

	offset := (params.Page - 1) * params.PageSize
	// id 辅助排序，防止大量 $sortField 重复导致乱序
	listSQL := fmt.Sprintf(
		"SELECT %s,COUNT() OVER() AS total_count FROM chuangliang_doris_cid.ad_order_pdd AS ad_act %s WHERE %s GROUP BY %s ORDER BY id ASC,%s DESC LIMIT %d OFFSET %d",
		selectFieldSQL, leftJoinSQL, advWhereSQL, groupByFieldSQL, sort, params.PageSize, offset,
	)

	e := data.NewDorisModel("", connects).QuerySQL(listSQL, &orders)
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	if len(orders) > 0 {
		total = orders[0].TotalCount

		var (
			emptyMainUserIdPid []string
			emptyNamePid       []string
		)
		for _, order := range orders {
			if order.MainUserId == 0 {
				emptyMainUserIdPid = append(emptyMainUserIdPid, order.PID)
			}
			if order.PName == "" {
				emptyNamePid = append(emptyNamePid, order.PID)
			}
		}

		type adminUser struct {
			mainUserId   int64
			userFullName string
		}
		pidMap := make(map[string]*adminUser)
		namePidMaps := make(map[string]string)
		if len(emptyMainUserIdPid) > 0 {
			emptyMainUserIdPid = util.ArrayUnique(emptyMainUserIdPid)
			var pidList []*common.PddPidList
			pidList, err = common.NewPddPidListModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("pid in ?", emptyMainUserIdPid)
			}, []string{"pid", "main_user_id"})
			if err != nil {
				return
			}
			MainUserIds := make([]int64, len(pidList))
			for i, list := range pidList {
				MainUserIds[i] = list.MainUserId
			}
			MainUserIds = util.ArrayUnique(MainUserIds)

			var userList []*base.User
			userList, err = base.NewUserModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("user_id in ?", MainUserIds)
			}, []string{"user_full_name", "user_id"})
			if err != nil {
				return
			}
			userMap := map[int64]string{}
			for _, user := range userList {
				userMap[user.UserId] = user.UserFullName
			}

			for _, list := range pidList {
				pidMap[list.Pid] = &adminUser{
					mainUserId:   list.MainUserId,
					userFullName: userMap[list.MainUserId],
				}
			}
		}
		if len(emptyNamePid) > 0 {
			emptyNamePid = util.ArrayUnique(emptyNamePid)
			var pidList []*common.PddPidList
			pidList, err = common.NewPddPidListModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("pid in ?", emptyNamePid)
			}, []string{"pid", "pid_name"})
			if err != nil {
				return
			}
			for _, list := range pidList {
				namePidMaps[list.Pid] = list.PidName
			}
		}

		// 商品备注
		for i, order := range orders {
			if typeDesc, ok := vars.PddOrderType[order.Type]; ok {
				orders[i].TypeDesc = typeDesc
			} else {
				orders[i].TypeDesc = "其他"
			}
			orderRefundTime, _ := order.OrderRefundTime.Value()
			if orderRefundTime != nil && orderRefundTime.(time.Time).Format(time.DateTime) == "1970-01-01 08:00:00" {
				parse, _ := time.Parse(time.DateTime, "0000-00-00 00:00:00")
				orders[i].OrderRefundTime = data.DbDateTime(parse)
			}
			if directDesc, ok := vars.OrderDirect[order.IsDirect]; ok {
				orders[i].IsDirectDesc = directDesc
			} else {
				orders[i].IsDirectDesc = "--"
			}
			//$v['trace_type_msg'] = PddOrderService::$traceTypeMsg[$v['cl_trace_type']] ? PddOrderService::$traceTypeMsg[$v['cl_trace_type']] : '--';
			//$v['csite_type_desc'] = OrderReportService::$csiteTypeDesc[$v['csite_type']]; //广告版位
			//$v['callback_event_desc'] = OrderReportService::$callbackEventDesc[$v['callback_event']]; //回传事件

			if order.MainUserId == 0 {
				if pid, ok := pidMap[order.PID]; ok {
					orders[i].MainUserId = pid.mainUserId
					orders[i].UserFullName = pid.userFullName
				}
			}
			if order.PName == "" {
				orders[i].PName = namePidMaps[order.PID]
			}
		}
	}

	return
}

func (p *PDD) getSortField(timeType int) (field string) {
	switch timeType {
	case TimeTypeCreateTime:
		field = "create_time"
		break
	case TimeTypeUpdateTime:
		field = "update_time"
	case TimeTypePayTime:
		field = p.ShopPayTimeField
		break
	case TimeTypeVerifyTime:
		field = p.VerifyTimeField
		break
	default:
		field = "order_create_time"
	}
	return
}
