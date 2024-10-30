package open_api

import (
	"errors"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type Order struct {
	C         *config.Config
	DbConnect *data.Data
}

var (
	ShopTypePDD  = vars.PlatformPdd
	ShopTypeJD   = vars.PlatformJd
	ShopTypeTB   = vars.PlatformTb
	OrderQueries = map[vars.Platform]OrderInterface{
		ShopTypePDD: &PDD{
			ShopPayTimeField: "order_pay_time",
			VerifyTimeField:  "order_verify_time",
		},
		ShopTypeJD: &JD{
			ShopPayTimeField: "payment_time",
			VerifyTimeField:  "payment_time",
		},
		ShopTypeTB: &TB{
			ShopPayTimeField: "tk_paid_time",
			VerifyTimeField:  "tk_paid_time",
		},
	}
)

const (
	TimeTypeCreateTime = iota + 1 // 创建时间
	TimeTypeUpdateTime            // 更新时间
	TimeTypePayTime               // 支付时间
	TimeTypeVerifyTime            // 审核时间
)

func (o *Order) OrderList(params statement.OrderList) ([]*OrderItem, int64, *errs.MyErr) {
	if fc, ok := OrderQueries[vars.Platform(params.ShopType)]; ok {
		return fc.GetOrderList(params, o.DbConnect)
	} else {
		return nil, 0, errs.Err(errs.SysError, errors.New("暂不支持该店铺类型"))
	}
}

type OrderInterface interface {
	GetOrderList(statement.OrderList, *data.Data) ([]*OrderItem, int64, *errs.MyErr)
}

type OrderItem struct {
	TotalCount      int64           `json:"total_count"`
	MainUserId      int64           `json:"main_user_id"`
	UserFullName    string          `json:"user_full_name"`
	PID             string          `json:"p_id"`
	PName           string          `json:"p_name"`
	Type            int             `json:"type"`
	TypeDesc        string          `json:"type_desc"`
	OrderRefundTime data.DbDateTime `json:"order_refund_time"`
	IsDirect        int             `json:"is_direct"`
	IsDirectDesc    string          `json:"is_direct_desc"`
}
