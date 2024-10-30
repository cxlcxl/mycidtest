package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type JD struct {
	ShopPayTimeField string
	VerifyTimeField  string
}

func (j *JD) GetOrderList(params statement.OrderList, connects *data.Data) (orders []*OrderItem, total int64, err *errs.MyErr) {
	return
}
