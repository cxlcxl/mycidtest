package base

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type UserCustom struct {
	GroupId        int64            `json:"group_id"`        //
	GroupName      string           `json:"group_name"`      // 权限组名
	Note           string           `json:"note"`            // 备注
	DefaultRoute   string           `json:"default_route"`   // 登录后默认路由
	MainUserId     int64            `json:"main_user_id"`    // 企业管理员ID
	Permissions    string           `json:"permissions"`     // 功能的权限，逗号隔开
	ModuleRange    string           `json:"module_range"`    // 模块的权限配置，json格式
	CreateTime     *data.DbDateTime `json:"create_time"`     // 创建时间
	CreateUserId   int64            `json:"create_user_id"`  // 创建用户ID
	UpdateTime     *data.DbDateTime `json:"update_time"`     // 最后更新时间
	UpdateUserId   int64            `json:"update_user_id"`  // 最后更新的用户ID
	IsDelete       uint8            `json:"is_delete"`       // 是否已删除
	ProductVersion uint8            `json:"product_version"` // 产品版本
}

type ModuleRange struct {
	InfoRange string `json:"info_range"`
	DataRange string `json:"data_range"`
}

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

func (m *UserCustomModel) FindByGroupId(groupId int64, fields []string) (userCustom *UserCustom, err *errs.MyErr) {
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
