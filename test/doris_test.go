package test

import (
	"testing"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/mylog"
)

type OrderPdd struct {
	MainUserId     int64
	AdId           int64
	Id             int64
	AdvertiserId   int64
	OrderAmount    float64
	GoodsName      string
	AdvertiserNick string
}

func TestDoris(t *testing.T) {
	sql := "select t0.main_user_id,t0.advertiser_id,t0.ad_id,t0.id,t0.order_amount,t0.goods_name,t1.advertiser_nick " +
		"from chuangliang_doris_cid.ad_order_pdd t0 " +
		"left join chuangliang_doris_common.media_account t1 on t0.main_user_id = t1.main_user_id and t0.advertiser_id = t1.advertiser_id " +
		"where t0.main_user_id = 12000020828"
	c, _ := config.LoadConfig("../config/config.yaml")
	log := mylog.NewLog()
	var tbOrders []OrderPdd
	err := data.NewDorisModel("", data.NewDB(c, log)).QuerySQL(sql, &tbOrders)
	t.Log(err, tbOrders)
}
