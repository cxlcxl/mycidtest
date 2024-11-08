package base

import (
	"errors"
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type MediaAccount struct {
	MediaAccountId        int64     `json:"media_account_id"`         // 媒体账号ID
	ParentId              int64     `json:"parent_id"`                // 父媒体账号id
	ToutiaoParentId       int64     `json:"toutiao_parent_id"`        // 从属头条的的客户中心id
	OwnerUserId           int64     `json:"owner_user_id"`            // 所属账户ID
	MediaType             string    `json:"media_type"`               // 媒体类型
	MediaAgentId          int64     `json:"media_agent_id"`           // 媒体代理ID
	AdvertiserType        int64     `json:"advertiser_type"`          // 媒体账号类型
	AdvertiserId          int64     `json:"advertiser_id"`            // 媒体平台的账号ID
	AdvertiserName        string    `json:"advertiser_name"`          // 媒体平台的账户名
	AdvertiserNick        string    `json:"advertiser_nick"`          // 媒体平台的账号昵称
	AdvertiserSource      string    `json:"advertiser_source"`        // 媒体账号来源
	AdvertiserStatus      int8      `json:"advertiser_status"`        // 媒体账号状态：1=正常，0=停投
	DevelopAppKey         string    `json:"develop_app_key"`          // 媒体平台开发者appkey
	DevelopAppSecret      string    `json:"develop_app_secret"`       // 媒体平台开发者appsecret
	AccessToken           string    `json:"access_token"`             // access token
	AccessTokenTime       time.Time `json:"access_token_time"`        // access token时间
	AccessTokenExpires    int64     `json:"access_token_expires"`     // access token过期时间
	AccessTokenRetryTimes int64     `json:"access_token_retry_times"` // access token重新刷新次数
	RefreshToken          string    `json:"refresh_token"`            // refresh token
	RefreshTokenExpires   string    `json:"refresh_token_expires"`    // refresh token过期时间
	Company               string    `json:"company"`                  // 公司名
	Note                  string    `json:"note"`                     // 备注
	CreateTime            time.Time `json:"create_time"`              // 创建时间
	CreateUserId          int64     `json:"create_user_id"`           // 创建用户ID
	UpdateTime            time.Time `json:"update_time"`              // 最后更新时间
	UpdateUserId          int64     `json:"update_user_id"`           // 最后更新的用户ID
	IsDelete              int8      `json:"is_delete"`                // 是否已删除
	Balance               float64   `json:"balance"`                  // 账号总余额(单位为：元)
	TodayCost             float64   `json:"today_cost"`               // 今日消耗(单位为：元)
	YesterdayCost         float64   `json:"yesterday_cost"`           // 昨日消耗(单位为：元)
	IsAutoClose           int8      `json:"is_auto_close"`            // 是否自动关闭，1开启，0不开启
	AiGray                string    `json:"ai_gray"`                  // 托管灰测接口
	AccountRole           string    `json:"account_role"`             // 新版授权账号角色
	ClAppId               int64     `json:"cl_app_id"`                // 创量应用的app_id
	MediaProjectId        int64     `json:"media_project_id"`         // 媒体项目ID
	MdmId                 int64     `json:"mdm_id"`                   // 广点通客户主体 id
	MdmName               string    `json:"mdm_name"`                 // 广点通客户主体名称
	IsActive              int8      `json:"is_active"`                // 是否活跃- 近30天有数据（包含当日）,账户级别总消耗大于0或- 30天内有新建计划的账户（包含当日），账户新建计划数大于0
	IsValidToken          int8      `json:"is_valid_token"`           // 1有效token2权限3token过期
	AuthorizeTime         time.Time `json:"authorize_time"`           // 授权时间
	AuthorizeType         int8      `json:"authorize_type"`           // 授权方式(0:默认 1:cookie)
	FirstIndustryName     string    `json:"first_industry_name"`      // 一级行业名
	SecondIndustryName    string    `json:"second_industry_name"`     // 二级行业名
	Rules                 string    `json:"rules"`                    // 规则
	IsShare               int8      `json:"is_share"`                 // 是否分享
	IsShow                int8      `json:"is_show"`                  // 是否显示
	Email                 string    `json:"email"`                    // 授权时登录的邮箱
	DisplayName           string    `json:"display_name"`             // 授权时登录用户名称
	MainUserId            int64     `json:"main_user_id"`             // 租户id
	DeductionBalance      float64   `json:"deduction_balance"`        // 消返红包余额
	GrantBalance          float64   `json:"grant_balance"`            // 赠款余额
	CashBalance           float64   `json:"cash_balance"`             // 非赠款余额
	EcpType               string    `json:"ecp_type"`                 // 账户类型
	CidAuthStatus         int8      `json:"cid_auth_status"`          // cid鉴权状态，0关闭，1开启
	Budget                float64   `json:"budget"`                   // 账号预算(单位为：元)
	ShareFromUserId       int64     `json:"share_from_user_id"`       // 分享者ID
	DeliveryStatus        int8      `json:"delivery_status"`          // 投放状态：0停投，1投放中，2未投放（cid新状态）
}

type MediaAccountModel struct {
	dbName string
	db     *gorm.DB
}

func NewMediaAccountModel(connect string, connects *data.Data) *MediaAccountModel {
	if connect == "" {
		connect = vars.DRUserMaster
	}
	return &MediaAccountModel{
		dbName: "media_account",
		db:     connects.DbConnects[connect],
	}
}

func (m *MediaAccountModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*MediaAccount, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = 0")
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

func (m *MediaAccountModel) GetOneByBuilder(builder data.QueryBuilder, fields []string) (one *MediaAccount, err *errs.MyErr) {
	query := m.db.Debug().Table(m.dbName)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Order("id desc").First(&one).Error
	if e != nil {
		if errors.Is(e, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errs.Err(errs.SysError, e)
	}
	return
}
