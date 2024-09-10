package order

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
)

type DorisOrder struct {
}

type DorisOrderModel struct {
	dbName string
	db     *gorm.DB
}

func NewDorisOrderModel(connect string, connects *data.Data) *DorisOrderModel {
	if connect == "" {
		connect = "doris_cid"
	}
	return &DorisOrderModel{
		dbName: "chuangliang_doris_cid",
		db:     connects.DbConnects[connect],
	}
}

func (m *DorisOrderModel) Query(sql string, value interface{}, parameters ...interface{}) (err error) {
	err = m.db.Raw(sql, parameters...).Scan(value).Error
	return
}
