package order

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/vars"
)

type TbOrder struct {
}

type TbOrderModel struct {
	dbName string
	db     *gorm.DB
}

func NewTbOrderModel(connect string, connects *data.Data) *TbOrderModel {
	if connect == "" {
		connect = vars.DROrderMaster
	}
	return &TbOrderModel{
		dbName: "ad_order_tb",
		db:     connects.DbConnects[connect],
	}
}
