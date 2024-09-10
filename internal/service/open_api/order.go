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
		ShopTypePDD: &PDD{},
	}
)

func (o *Order) OrderList(params statement.OrderList) ([]*OrderItem, int64, *errs.MyErr) {
	if fc, ok := OrderQueries[vars.Platform(params.ShopType)]; ok {
		return fc.GetOrderList(params)
	} else {
		return nil, 0, errs.Err(errs.SysError, errors.New("暂不支持该店铺类型"))
	}
}

type OrderInterface interface {
	GetOrderList(statement.OrderList) ([]*OrderItem, int64, *errs.MyErr)
}

type OrderItem struct {
}
