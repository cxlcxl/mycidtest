package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type UserVersion struct {
	ID               int64            `json:"id"`
	UserId           int64            `json:"user_id"`
	ProductVersion   int8             `json:"product_version"`
	QianChuanVersion int8             `json:"qianchuan_version"`
	GroupId          int64            `json:"group_id"`
	EndDate          *data.DbDate     `json:"end_date"`
	IsDelete         uint8            `json:"is_delete"`
	CreateTime       *data.DbDateTime `json:"create_time"`  // 创建时间
	UpdateTime       *data.DbDateTime `json:"update_time"`  //
	MainUserId       int64            `json:"main_user_id"` // 租户id
}

type UserVersionModel struct {
	dbName string
	db     *gorm.DB
}

func NewUserVersionModel(connect string, connects *data.Data) *UserVersionModel {
	if connect == "" {
		connect = vars.DRActCLAdCommon
	}
	return &UserVersionModel{
		dbName: "admin_user_zone",
		db:     connects.DbConnects[connect],
	}
}

func (m *UserVersionModel) GetAdminUserVersionInfoByBuilder(builder data.QueryBuilder, fields []string) (version *UserVersion, err *errs.MyErr) {
	query := m.db.Table(m.dbName)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.First(&version).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}

func (m *UserVersionModel) GetAdminUserVersionListByBuilder(builder data.QueryBuilder, fields []string) (versions []*UserVersion, err *errs.MyErr) {
	query := m.db.Table(m.dbName)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&versions).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}
