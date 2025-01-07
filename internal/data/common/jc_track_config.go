package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type JcTrackConfig struct {
	ID                     int64  `json:"id"`
	Name                   string `json:"name"`
	Pid                    string `json:"pid"`
	MainUserId             int64  `json:"main_user_id"`        // 所属的用户id
	MediaType              int8   `json:"media_type"`          // 广告平台
	CallbackEventType      string `json:"callback_event_type"` // 回调事件
	CallbackAccount        string `json:"callback_account"`    // 回调账号
	PartnerId              string `json:"partner_id"`
	RequestFrom            string `json:"request_from"`
	Benefit                string `json:"benefit"`
	PartnerMin             int64  `json:"partner_min"` // partner数字范围最小值
	PartnerMax             int64  `json:"partner_max"`
	CallbackTransformTypes string `json:"callback_transform_types"` // 配置的回传类型
}

// TableName 表名称
func (*JcTrackConfig) TableName() string {
	return "jc_track_config"
}

type JcTrackConfigModel struct {
	dbName string
	db     *gorm.DB
}

func NewJcTrackConfigModel(connect string, connects *data.Data) *JcTrackConfigModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &JcTrackConfigModel{
		dbName: "jc_track_config",
		db:     connects.DbConnects[connect],
	}
}

func (m *JcTrackConfigModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*JcTrackConfig, err *errs.MyErr) {
	query := m.db.Debug().Table(m.dbName)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&list).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}
