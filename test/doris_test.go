package test

import (
	"testing"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/mylog"
)

type OrderPdd struct {
	MainUserId   int64
	AdId         int64
	Id           int64
	AdvertiserId int64
	OrderAmount  float64
	GoodsName    string
}

func TestDoris(t *testing.T) {
	sql := "select * from chuangliang_doris_cid.ad_order_pdd where main_user_id = 12000020828 limit 1"
	c, _ := config.LoadConfig("../config/config.yaml")
	log := mylog.NewLog()
	var tbOrders OrderPdd
	err := data.NewDorisModel("", data.NewDB(c, log)).QuerySQL(sql, &tbOrders)
	t.Log(err, tbOrders)
}
