package vars

const (
	PddOrderType0   = 0
	PddOrderType1   = 1
	PddOrderType2   = 2
	PddOrderType3   = 4
	PddOrderType6   = 6
	PddOrderType7   = 7
	PddOrderType8   = 8
	PddOrderType10  = 10
	PddOrderType16  = 16
	PddOrderType51  = 51
	PddOrderType64  = 64
	PddOrderType77  = 77
	PddOrderType94  = 94
	PddOrderType101 = 101
	PddOrderType103 = 103
	PddOrderType104 = 104
	PddOrderType105 = 105

	OrderNotDirect = 0
	OrderIsDirect  = 1

	TraceTypeSameStore  = 2
	TraceTypeCrossStore = 3
)
const (
	AdSiteTypeDefault = iota
)
const (
	CallbackEventDefault = iota
)

var (
	OrderDirect = map[int]string{
		OrderNotDirect: "非直推订单",
		OrderIsDirect:  "直推订单",
	}
	TraceType = map[int]string{
		TraceTypeSameStore:  "同店",
		TraceTypeCrossStore: "跨店",
	}
	PddOrderType = map[int]string{
		PddOrderType0:   "单品",
		PddOrderType1:   "红包活动推广",
		PddOrderType2:   "领券页推荐",
		PddOrderType3:   "多多进宝商城推广",
		PddOrderType6:   "拼团后推荐",
		PddOrderType7:   "今日爆款",
		PddOrderType8:   "品牌清仓",
		PddOrderType10:  "全店关联",
		PddOrderType16:  "支付新用户锁佣",
		PddOrderType51:  "商详推荐",
		PddOrderType64:  "跨店关联",
		PddOrderType77:  "刮刮卡活动推广",
		PddOrderType94:  "充值中心",
		PddOrderType101: "品牌黑卡",
		PddOrderType103: "百亿补贴频道",
		PddOrderType104: "内购清单频道",
		PddOrderType105: "超级红包",
	}
	// AdSiteType 广告版位映射关系
	AdSiteType = map[int]string{
		0: "--",
		1: "抖音",
		2: "今日头条",
		3: "西瓜视频",
		4: "火山小视频",
		5: "番茄小说",
		6: "穿山甲开屏广告",
		7: "穿山甲网盟非开屏广告",
		8: "通投广告位",
		9: "搜索",
	}
	// CallbackEvent 回传事件
	CallbackEvent = map[int]string{
		0: "--",
		1: "APP内下单",
		2: "付费",
		3: "APP内下单、付费",
	}
)
