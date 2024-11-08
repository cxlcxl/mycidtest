package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type PddGoodsNoteRel struct {
	ID          int64  `json:"id"`
	RecordId    int64  `json:"record_id"`     // 记录id
	GoodsId     int64  `json:"goods_id"`      // 商品id
	Note        string `json:"note"`          // 备注
	MainUserId  int64  `json:"main_user_id"`  // 所属的用户id
	OwnerUserId int64  `json:"owner_user_id"` // 用户id
	IsDelete    int8   `json:"is_delete"`
	*data.Timestamp
}

type PddGoodsNoteRelModel struct {
	dbName string
	db     *gorm.DB
}

func NewPddGoodsNoteRelModel(connect string, connects *data.Data) *PddGoodsNoteRelModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &PddGoodsNoteRelModel{
		dbName: "pdd_goods_note_rel",
		db:     connects.DbConnects[connect],
	}
}

func (m *PddGoodsNoteRelModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*PddGoodsNoteRel, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = ?", 0)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&list).Error
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	return
}
