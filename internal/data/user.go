package data

import (
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/pkg/errs"
)

type User struct {
	UserId                 int64     `json:"user_id" gorm:"user_id"`                                     // 用户id
	ProjectId              int64     `json:"project_id" gorm:"project_id"`                               // 项目组ID
	GroupId                int64     `json:"group_id" gorm:"group_id"`                                   // 权限组ID
	Email                  string    `json:"email" gorm:"email"`                                         // 邮箱
	UserName               string    `json:"user_name" gorm:"user_name"`                                 // 用户名
	UserFullName           string    `json:"user_full_name" gorm:"user_full_name"`                       // 用户全称
	Password               string    `json:"-" gorm:"password"`                                          // 密码
	Mobile                 string    `json:"mobile" gorm:"mobile"`                                       // 手机号码
	Note                   string    `json:"note" gorm:"note"`                                           // 备注
	IsLock                 int8      `json:"is_lock" gorm:"is_lock"`                                     // 是否锁定
	ParentId               int64     `json:"parent_id" gorm:"parent_id"`                                 // 所属用户id
	CreateTime             time.Time `json:"create_time" gorm:"create_time"`                             // 创建时间
	CreateUserId           int64     `json:"create_user_id" gorm:"create_user_id"`                       // 创建用户ID
	UpdateTime             time.Time `json:"update_time" gorm:"update_time"`                             // 最后更新时间
	UpdateUserId           int64     `json:"update_user_id" gorm:"update_user_id"`                       // 最后更新的用户ID
	GrayPermissions        string    `json:"gray_permissions" gorm:"gray_permissions"`                   // 灰发权限，逗号隔开
	DataRange              string    `json:"data_range" gorm:"data_range"`                               // 数据权限，默认是从权限组继承
	IsDelete               int8      `json:"is_delete" gorm:"is_delete"`                                 // 是否已删除
	QqLoginOpenid          string    `json:"qq_login_openid" gorm:"qq_login_openid"`                     // qq登录对应的openid
	DevelopKey             string    `json:"develop_key" gorm:"develop_key"`                             // 开发者key，对外api需要在主账号配置
	DevelopIpWhitelist     string    `json:"develop_ip_whitelist" gorm:"develop_ip_whitelist"`           //
	Contact                string    `json:"contact" gorm:"contact"`                                     // 联系人
	SignStatus             int8      `json:"sign_status" gorm:"sign_status"`                             // 签约状态 1.未签约 2.已签约 3.解约
	IsMonitor              int8      `json:"is_monitor" gorm:"is_monitor"`                               // 是否监控 0：否，1：是
	SignTime               time.Time `json:"sign_time" gorm:"sign_time"`                                 // 签约时间
	EndDate                time.Time `json:"end_date" gorm:"end_date"`                                   // 到期日期
	Level                  int8      `json:"level" gorm:"level"`                                         // 客户等级 1：S级，2：A级，3：B级，4：C级
	Operator               string    `json:"operator" gorm:"operator"`                                   // 运营负责人
	BusinessMember         string    `json:"business_member" gorm:"business_member"`                     // 商务责任人
	LatestNews             string    `json:"latest_news" gorm:"latest_news"`                             // 最新动态
	CompanyType            int8      `json:"company_type" gorm:"company_type"`                           // 企业类型 1.广告主 2.代理
	Industry               string    `json:"industry" gorm:"industry"`                                   // 行业
	MediaLaunch            string    `json:"media_launch" gorm:"media_launch"`                           // 投放媒体
	CreateLevel            string    `json:"create_level" gorm:"create_level"`                           // 新建广告的等级
	UsedStatus             int64     `json:"used_status" gorm:"used_status"`                             // 使用状态
	SignProgress           int64     `json:"sign_progress" gorm:"sign_progress"`                         // 签约进度
	AgreementEndDate       time.Time `json:"agreement_end_date" gorm:"agreement_end_date"`               // 合同结束日期
	AgreementStartDate     time.Time `json:"agreement_start_date" gorm:"agreement_start_date"`           // 合同开始日期
	Competition            string    `json:"competition" gorm:"competition"`                             // 竞品
	MainUserId             int64     `json:"main_user_id" gorm:"main_user_id"`                           // 租户id
	Target                 string    `json:"target" gorm:"target"`                                       // 投放目标
	LastChangePasswordTime time.Time `json:"last_change_password_time" gorm:"last_change_password_time"` // 最后一次修改密码时间
	ContractType           int8      `json:"contract_type" gorm:"contract_type"`                         // 新老合同 0：老，1：新
	ExternalUserId         int64     `json:"external_user_id" gorm:"external_user_id"`                   // 关联的外部用户ID
	ActiveDate             time.Time `json:"active_date" gorm:"active_date"`                             // 最近活跃时间
	BetaVersion            int8      `json:"beta_version" gorm:"beta_version"`                           // 标记beta版本（0-非beta版本,1-beta版本）
	ChargePersonType       int8      `json:"charge_person_type" gorm:"charge_person_type"`               // 负责人类型，0-无，1-商务，2-代理商，3-自营负责人
	ChargePersonId         int64     `json:"charge_person_id" gorm:"charge_person_id"`                   // 负责人用户ID
	City                   string    `json:"city" gorm:"city"`                                           // 城市
	AgencyOperations       string    `json:"agency_operations" gorm:"agency_operations"`                 // 代运营
	IsAgencyOperations     int8      `json:"is_agency_operations" gorm:"is_agency_operations"`           // 是否代运营0-否，1-是
	Ae                     string    `json:"ae" gorm:"ae"`                                               // AE
}

type UserModel struct {
	dbName string
	db     *gorm.DB
}

func NewUserModel(connect string, connects *Data) *UserModel {
	if connect == "" {
		connect = "user_master"
	}
	return &UserModel{
		dbName: "admin_user",
		db:     connects.DbConnects[connect],
	}
}

func (m *UserModel) FindUserByLogin(email, password string) (user *User, err *errs.MyErr) {
	e := m.db.Table(m.dbName).
		Where("email = ? and password = ?", email, password).
		Where("is_delete = ?", 0).
		First(&User{}).Error
	if e != nil {
		return nil, errs.Err(errs.LoginFinUserError, e)
	}
	return
}
