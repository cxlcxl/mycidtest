package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/pkg/errs"
)

type PDD struct {
}

func (p *PDD) GetOrderList(params statement.OrderList) (orders []*OrderItem, total int64, err *errs.MyErr) {
	return
}
