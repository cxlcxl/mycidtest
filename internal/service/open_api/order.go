package open_api

import (
	"errors"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type Order struct {
	DbConnect *data.Data
}

var (
	ShopTypePDD  = vars.PlatformPdd
	ShopTypeJD   = vars.PlatformJd
	ShopTypeTB   = vars.PlatformTb
	OrderQueries = map[vars.Platform]OrderInterface{
		ShopTypePDD: &PDD{
			PayTimeField:    "order_pay_time",
			VerifyTimeField: "order_verify_time",
		},
		ShopTypeJD: &JD{
			PayTimeField:    "payment_time",
			VerifyTimeField: "payment_time",
		},
		ShopTypeTB: &TB{
			PayTimeField:    "tk_paid_time",
			VerifyTimeField: "tk_paid_time",
		},
	}
)

const (
	TimeTypeCreateTime = iota + 1 // 创建时间
	TimeTypeUpdateTime            // 更新时间
	TimeTypePayTime               // 支付时间
	TimeTypeVerifyTime            // 审核时间
)

func (o *Order) OrderList(params statement.OrderList) (interface{}, int64, *errs.MyErr) {
	if fc, ok := OrderQueries[vars.Platform(params.ShopType)]; ok {
		return fc.GetOrderList(params, o.DbConnect)
	} else {
		return nil, 0, errs.Err(errs.SysError, errors.New("暂不支持该店铺类型"))
	}
}

type OrderInterface interface {
	GetOrderList(statement.OrderList, *data.Data) (interface{}, int64, *errs.MyErr)
}

type PddOrderItem struct {
	Id                int64   `json:"id"`
	OrderSn           string  `json:"order_sn"`
	MainUserId        int64   `json:"main_user_id"`
	OwnerUserId       int64   `json:"owner_user_id"`
	AdvertiserNick    string  `json:"advertiser_nick"`
	UserName          string  `json:"user_name"`
	UserFullName      string  `json:"user_full_name"`
	PID               string  `json:"p_id"`
	PName             string  `json:"p_name"`
	IsHide            int     `json:"is_hide"`
	GoodsId           int64   `json:"goods_id"`
	GoodsName         string  `json:"goods_name"`
	GoodsQuantity     int     `json:"goods_quantity"`
	PromotionRate     int     `json:"promotion_rate"`
	GoodsThumbnailUrl string  `json:"goods_thumbnail_url"`
	FailReason        string  `json:"fail_reason"`
	MallId            int64   `json:"mall_id"`
	MallName          string  `json:"mall_name"`
	OrderStatus       int     `json:"order_status"`
	OrderStatusDesc   string  `json:"order_status_desc"`
	OrderAmount       float64 `json:"order_amount"`
	Type              int     `json:"type"`
	ClickId           int64   `json:"click_id"`
	AdId              int64   `json:"ad_id"`
	AdSiteId          int64   `json:"ad_site_id"`
	MediaType         int     `json:"media_type"`
	AccountUserId     int64   `json:"account_user_id"`
	IsCallback        int     `json:"is_callback"`
	NoCallbackReason  string  `json:"no_callback_reason"`
	CallbackType      int     `json:"callback_type"`
	TypeDesc          string  `json:"type_desc"`
	OrderRefundTime   string  `json:"order_refund_time"`
	OrderPayTime      string  `json:"order_pay_time"`
	OrderVerifyTime   string  `json:"order_verify_time"`
	OrderCreateTime   string  `json:"order_create_time"`
	OrderReceiveTime  string  `json:"order_receive_time"`
	CallbackTime      string  `json:"callback_time"`
	PromotionType     int     `json:"promotion_type"`
	IsDiffGoods       int     `json:"is_diff_goods"`
	CustomParameters  string  `json:"custom_parameters"`
	IsDirect          int     `json:"is_direct"`
	IsDirectDesc      string  `json:"is_direct_desc"`
	ClTraceType       int     `json:"cl_trace_type"`
	TraceTypeMsg      string  `json:"trace_type_msg"`
	CallbackEvent     int     `json:"callback_event"`
	CallbackEventDesc string  `json:"callback_event_desc"`
	CSiteType         int     `json:"csite_type"`
	CSiteTypeDesc     string  `json:"csite_type_desc"`
	PromotionAmount   float64 `json:"promotion_amount"`
	GoodsPrice        float64 `json:"goods_price"`
	CreateTime        string  `json:"create_time"`
	UpdateTime        string  `json:"update_time"`
	TotalCount        int64   `json:"total_count"`
}
