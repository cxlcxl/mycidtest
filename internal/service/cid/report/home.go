package report

import (
	"fmt"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type HomeService struct {
	C         *config.Config
	DbConnect *data.Data
}

type OrderSumItem struct {
	OrderCnt          int     `json:"order_cnt"`
	OrderAmount       float64 `json:"order_amount"`
	AverageOrderPrice float64 `json:"average_order_price"`
	OrderRefundNum    int     `json:"order_refund_num"`
	OrderRefundAmount float64 `json:"order_refund_amount"`
	TotalFee          float64 `json:"total_fee"`
	OrderRefundRate   float64 `json:"order_refund_rate"`
	TradedOrderCount  int     `json:"traded_order_count"`
	TradedOrderAmount float64 `json:"traded_order_amount"`
}

func (s *HomeService) OrderSum(params statement.ReportHomeOrderSum) (orders []*OrderSumItem, err *errs.MyErr) {
	sql := ""
	var dd interface{}
	e := data.NewDorisModel("", s.DbConnect).QuerySQL(sql, &dd)
	fmt.Println(e)
	return
}
