package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type JcLink struct {
	ID            int64  `json:"id"`              // 链接ID，其他关联功能表中的 ad_site_id
	MainUserId    int64  `json:"main_user_id"`    // 所属的用户id
	MediaType     int8   `json:"media_type"`      // 广告平台
	ExposeLink    string `json:"expose_link"`     // 曝光检测链接
	ClickLink     string `json:"click_link"`      // 点击检测链接
	Remark        string `json:"remark"`          // 备注
	CreateUserId  int64  `json:"create_user_id"`  //
	IsDelete      int8   `json:"is_delete"`       //
	ExposeLinkMd5 string `json:"expose_link_md5"` // 曝光检测链接md5
	ClickLinkMd5  string `json:"click_link_md5"`  // 点击检测链接md5
	JcConfigId    int64  `json:"jc_config_id"`    // 取哪套回传配置
	*data.Timestamp
}

// TableName 表名称
func (*JcLink) TableName() string {
	return "jc_link"
}

type JcLinkModel struct {
	dbName string
	db     *gorm.DB
}

func NewJcLinkModel(connect string, connects *data.Data) *JcLinkModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &JcLinkModel{
		dbName: "jc_link",
		db:     connects.DbConnects[connect],
	}
}

func (m *JcLinkModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*JcLink, err *errs.MyErr) {
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
