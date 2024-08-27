package auth_token

import "github.com/golang-jwt/jwt/v5"

type LoginData struct {
	UserId         int64  `json:"user_id"`         // 用户id
	ProjectId      int64  `json:"project_id"`      // 项目组ID
	GroupId        int64  `json:"group_id"`        // 权限组ID
	Email          string `json:"email"`           // 邮箱
	UserName       string `json:"user_name"`       // 用户名
	UserFullName   string `json:"user_full_name"`  // 用户全称
	Mobile         string `json:"mobile"`          // 手机号码
	DataRange      string `json:"data_range"`      // 数据权限，默认是从权限组继承
	LatestNews     string `json:"latest_news"`     // 最新动态
	CompanyType    int8   `json:"company_type"`    // 企业类型 1.广告主 2.代理
	Industry       string `json:"industry"`        // 行业
	MediaLaunch    string `json:"media_launch"`    // 投放媒体
	CreateLevel    string `json:"create_level"`    // 新建广告的等级
	UsedStatus     int64  `json:"used_status"`     // 使用状态
	MainUserId     int64  `json:"main_user_id"`    // 租户id
	ContractType   int8   `json:"contract_type"`   // 新老合同 0：老，1：新
	ProductVersion int    `json:"product_version"` // 产品版本
}

type LoginToken struct {
	UserInfo *LoginData `json:"user_info"`
	Token    TokenInfo  `json:"token"`
}

type TokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpireTime  int64  `json:"expire_time"`
}

type LoginClaims struct {
	UserInfo *LoginData `json:"user_info"`
	jwt.RegisteredClaims
}
