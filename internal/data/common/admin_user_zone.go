package data

import (
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type UserZone struct {
	ID         int64     `json:"id"`
	Email      string    `json:"email"`        // 邮箱
	ZoneIndex  int       `json:"zone_index"`   // 分区编号
	CreateTime time.Time `json:"create_time"`  // 创建时间
	MainUserId int64     `json:"main_user_id"` // 租户id
}

type UserZoneModel struct {
	dbName string
	db     *gorm.DB
}

func NewUserZoneModel(connect string, connects *data.Data) *UserZoneModel {
	if connect == "" {
		connect = "common"
	}
	return &UserZoneModel{
		dbName: "admin_user_zone",
		db:     connects.DbConnects[connect],
	}
}

func (m *UserZoneModel) FindUserZoneByEmail(email string) (userZone *UserZone, err *errs.MyErr) {
	e := m.db.Table(m.dbName).Where("email = ?", email).First(&userZone).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}

	return
}
