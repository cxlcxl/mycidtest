package vars

import (
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/mylog"
)

type Platform int
type Media int

var (
	// BasePath 系统跟路径
	BasePath string
	SysLog   *mylog.Log
	Config   *config.Config
)

const (
	LoginKey        = "__sys_login_key__"
	OpenApiLoginKey = "__sys_api_login_key__"
	RequestIdKey    = "__sys_request_id_key__"

	PlatformTb  Platform = 1
	PlatformPdd Platform = 2
	PlatformJd  Platform = 3

	MediaTypeIntTT  Media = 1
	MediaTypeIntKS  Media = 2
	MediaTypeIntGDT Media = 3

	MediaTypeStringTT  = "toutiao"
	MediaTypeStringKS  = "kuaishou"
	MediaTypeStringGDT = "gdt"

	NotifyTypeGoodsReport                = "pdd_goods_report"          // 商品报备通知
	NotifyTypeGoodsPromotionLimit        = "pdd_goods_promotion_limit" // 商品推广限制通知
	NotifyTypeRefundRate                 = "pdd_refund_rate"           // 商品退单率过高通知
	NotifyTypeCallbackRate               = "callback_rate"             // 账户回传率误差通知
	NotifyTypeAdBind                     = "ad_bind"                   // 绑定广告账户通知
	NotifyTypeAccountAlarm               = "account_alarm"             //账户预警
	NotifyTypeRiskStrategyNoReplaceGoods = "risk_strategy_no_replace_goods"

	NotifyMethodFeiShu   = "feishu_robot"
	NotifyMethodWechat   = "qiyeweixin_robot"
	NotifyMethodDingDing = "dingding_robot"
	NotifyMethodSMS      = "sms"
)

const (
	NotifyLevelCompany = iota + 1
	NotifyLevelDepartment
	NotifyLevelPerson
)

var (
	MediaTypeInt2Str = map[Media]string{
		MediaTypeIntTT:  MediaTypeStringTT,
		MediaTypeIntKS:  MediaTypeStringKS,
		MediaTypeIntGDT: MediaTypeStringGDT,
	}
	NotifyLevel = map[int]string{
		NotifyLevelCompany:    "企业",
		NotifyLevelDepartment: "部门",
		NotifyLevelPerson:     "个人",
	}
	// NotifyType 通知类型
	NotifyType = []string{
		NotifyTypeGoodsReport,
		NotifyTypeGoodsPromotionLimit,
		NotifyTypeRefundRate,
		NotifyTypeCallbackRate,
		NotifyTypeAdBind,
		NotifyTypeAccountAlarm,
		NotifyTypeRiskStrategyNoReplaceGoods,
	}
	// NotifyMethod 通知方式
	NotifyMethod = []string{
		NotifyMethodFeiShu,
		NotifyMethodWechat,
		NotifyMethodDingDing,
		NotifyMethodSMS,
	}
	// NotifyTypeGoodsPromotionLimitValue 推广异常模板key对应的设置值
	NotifyTypeGoodsPromotionLimitValue = map[string]int{
		"promotionLimitNotifyTemplate": 1,
		"promotionNotifyTemplate":      2,
	}
)
