package open_api

import (
	"errors"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type Order struct {
	C         *config.Config
	DbConnect *data.Data
}

var (
	ShopTypePDD  = 1
	ShopTypeJD   = 2
	ShopTypeTB   = 3
	OrderQueries = map[int]OrderInterface{
		ShopTypePDD: &PDD{},
	}
)

func (o *Order) OrderList(params statement.OrderList) ([]*OrderItem, int64, *errs.MyErr) {
	if fc, ok := OrderQueries[params.ShopType]; ok {
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
