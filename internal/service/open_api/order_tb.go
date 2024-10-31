package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type TB struct {
	PayTimeField    string
	VerifyTimeField string
}

func (t *TB) GetOrderList(params statement.OrderList, connects *data.Data) (orders interface{}, total int64, err *errs.MyErr) {
	return
}
