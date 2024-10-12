package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type UserCustomModel struct {
	dbName string
	db     *gorm.DB
}

func NewUserCustomModel(connect string, connects *data.Data) *UserCustomModel {
	if connect == "" {
		connect = vars.DRActCLAd
	}
	return &UserCustomModel{
		dbName: "admin_user_custom",
		db:     connects.DbConnects[connect],
	}
}

func (m *UserCustomModel) FindByGroupId(groupId int64, fields []string) (userCustom *base.UserCustom, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("group_id = ?", groupId).Where("is_delete = 0")
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	e := query.First(&userCustom).Error
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	return
}
