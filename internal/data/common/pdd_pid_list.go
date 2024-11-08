package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type PddPidList struct {
	ID            int64  `json:"id"`
	MemberId      int64  `json:"member_id"`       // 拼多多会员id
	AccountUserId int64  `json:"account_user_id"` //
	PidName       string `json:"pid_name"`        // 推广位名称
	Pid           string `json:"pid"`             // 推广位
	AdZoneId      int64  `json:"ad_zone_id"`      //
	OAuthStatus   int8   `json:"oauth_status"`    // 推广位授权状态
	GoodsSign     string `json:"goods_sign"`      // 拼多多商品goodsSign
	MediaId       int64  `json:"media_id"`        // 绑定的媒体备案id
	Status        uint8  `json:"status"`          // 推广位状态：0-正常，1-封禁
	Note          string `json:"note"`
	MainUserId    int64  `json:"main_user_id"`
	OwnerUserId   int64  `json:"owner_user_id"`
	IsDelete      uint8  `json:"is_delete"`
	*data.Timestamp
}

type PddPidListModel struct {
	dbName string
	db     *gorm.DB
}

func NewPddPidListModel(connect string, connects *data.Data) *PddPidListModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &PddPidListModel{
		dbName: "pdd_pid_list",
		db:     connects.DbConnects[connect],
	}
}

func (m *PddPidListModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*PddPidList, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = ?", 0)
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
