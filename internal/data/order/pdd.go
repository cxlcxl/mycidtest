package order

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/vars"
)

type PddOrder struct {
}

type PddOrderModel struct {
	dbName string
	db     *gorm.DB
}

func NewPddOrderModel(connect string, connects *data.Data) *PddOrderModel {
	if connect == "" {
		connect = vars.DROrderMaster
	}
	return &PddOrderModel{
		dbName: "ad_order_pdd",
		db:     connects.DbConnects[connect],
	}
}
