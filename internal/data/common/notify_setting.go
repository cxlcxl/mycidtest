package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type NotifySetting struct {
	ID            int64  `json:"id"`
	NotifyLevel   int8   `json:"notify_level"`
	OwnerUserId   int64  `json:"owner_user_id"`
	ExtendedValue string `json:"extended_value"`
	Webhook       string `json:"webhook"`
	NotifyMethod  string `json:"notify_method"`
}

type NotifySettingModel struct {
	dbName string
	db     *gorm.DB
}

func NewNotifySettingModel(connect string, connects *data.Data) *NotifySettingModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &NotifySettingModel{
		dbName: "notify_setting",
		db:     connects.DbConnects[connect],
	}
}

func (m *NotifySettingModel) GetNotifySettingListByBuilder(builder data.QueryBuilder, fields []string) (notifySettings []*NotifySetting, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = ?", 0)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&notifySettings).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}
