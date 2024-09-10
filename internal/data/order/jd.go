package order

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/vars"
)

type JdOrder struct {
}

type JdOrderModel struct {
	dbName string
	db     *gorm.DB
}

func NewJdOrderModel(connect string, connects *data.Data) *JdOrderModel {
	if connect == "" {
		connect = vars.DROrderMaster
	}
	return &JdOrderModel{
		dbName: "ad_order_jd",
		db:     connects.DbConnects[connect],
	}
}
