package base

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type AdminProject struct {
	ProjectId    int64            `json:"project_id"`     //
	ProjectName  string           `json:"project_name"`   // 项目组名字
	OwnerUserId  int64            `json:"owner_user_id"`  // 所属主账号id
	Note         string           `json:"note"`           //
	CreateTime   *data.DbDateTime `json:"create_time"`    // 创建时间
	CreateUserId int64            `json:"create_user_id"` // 创建用户ID
	UpdateTime   *data.DbDateTime `json:"update_time"`    // 最后更新时间
	UpdateUserId int64            `json:"update_user_id"` // 最后更新的用户ID
	IsDelete     uint8            `json:"is_delete"`      // 是否已删除
	PathLevel    string           `json:"path_level"`     // 部门全层级路径json
	ParentId     int64            `json:"parent_id"`      //
	MainUserId   int64            `json:"main_user_id"`   //
}

type AdminProjectModel struct {
	dbName string
	db     *gorm.DB
}

func NewAdminProjectModel(connect string, connects *data.Data) *AdminProjectModel {
	if connect == "" {
		connect = vars.DRActCLAd
	}
	return &AdminProjectModel{
		dbName: "admin_project",
		db:     connects.DbConnects[connect],
	}
}

func (m *AdminProjectModel) FindByOwnerUserId(ownerUserId int64, fields []string) (projects []*AdminProject, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("owner_user_id = ?", ownerUserId).Where("is_delete = 0")
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	e := query.Find(&projects).Error
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	return
}
