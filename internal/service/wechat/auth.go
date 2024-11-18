package wechat

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"net/http"
	"xiaoniuds.com/cid/app/wechat/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type Service struct {
	DbConnect *data.Data
}

type MiniProgramLoginResponse struct {
	OpenId     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrMsg     string `json:"errmsg"`
}

type MiniProgramUser struct {
	UserId       int64  `json:"user_id"`
	UserName     string `json:"user_name"`
	UserFullName string `json:"user_full_name"`
	Email        string `json:"email"`
	MainUserId   int64  `json:"main_user_id"`
}

func (s *Service) GetUsersByCode(code string) (users []*MiniProgramUser, openId string, err *errs.MyErr) {
	response, e := resty.New().R().SetQueryParams(map[string]string{
		"appid":      vars.MiniProgramAppId,
		"secret":     vars.MiniProgramAppSecret,
		"js_code":    code,
		"grant_type": "authorization_code",
	}).Get(vars.MiniProgramHostCode2Session)
	if e != nil {
		return nil, "", errs.Err(errs.WechatErrRequest, e)
	}

	if response.StatusCode() != http.StatusOK {
		return nil, "", errs.Err(errs.WechatErrRequest, fmt.Errorf(response.Status()))
	}

	var loginInfo MiniProgramLoginResponse
	e = json.Unmarshal(response.Body(), &loginInfo)
	if e != nil {
		return nil, "", errs.Err(errs.ErrJsonUnmarshal, e)
	}
	if loginInfo.OpenId == "" {
		return nil, "", errs.Err(errs.WechatErrRequest, fmt.Errorf("[openid is empty] %s", loginInfo.ErrMsg))
	}

	openId = loginInfo.OpenId
	adminUsers, err := base.NewUserModel("cid_test", s.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_delete = 0 and is_lock = 0 and openid = ?", loginInfo.OpenId)
	}, []string{"user_id", "user_name", "user_full_name", "email", "main_user_id"})
	if err == nil {
		_ = copier.Copy(&users, &adminUsers)
	}
	return
}

func (s *Service) BindUser(bindData statement.BindUser) (users []*MiniProgramUser, err *errs.MyErr) {
	user, err := base.NewUserModel("cid_test", s.DbConnect).FindUserByLogin(bindData.Email, util.Password(bindData.Pass, false))
	if err != nil {
		return
	}
	err = base.NewUserModel("cid_test", s.DbConnect).UpdateByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("user_id = ?", user.UserId)
	}, map[string]interface{}{"openid": bindData.OpenId})
	if err != nil {
		return
	}

	adminUsers, err := base.NewUserModel("cid_test", s.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_delete = 0 and is_lock = 0 and openid = ?", bindData.OpenId)
	}, []string{"user_id", "user_name", "user_full_name", "email", "main_user_id"})
	if err == nil {
		_ = copier.Copy(&users, &adminUsers)
	}
	return
}

func (s *Service) SelectUserLogin(openid string, userId int64) (user *MiniProgramUser, token string, err *errs.MyErr) {
	adminUser, err := base.NewUserModel("cid_test", s.DbConnect).FindUserById(userId)
	if err != nil {
		return
	}
	if adminUser.OpenId != openid {
		return nil, "", errs.Err(errs.OpenApiErrUserNotMatch, nil)
	}
	_ = copier.Copy(&user, &adminUser)

	tokenBuilder := &auth_token.WechatMiniProgramToken{
		Data: &auth_token.LoginData{
			UserId:         adminUser.UserId,
			ProjectId:      adminUser.ProjectId,
			GroupId:        adminUser.GroupId,
			GroupName:      adminUser.UserName,
			Email:          adminUser.Email,
			UserName:       adminUser.UserName,
			UserFullName:   adminUser.UserFullName,
			Mobile:         adminUser.Mobile,
			DataRange:      adminUser.DataRange,
			CompanyType:    adminUser.CompanyType,
			Industry:       adminUser.Industry,
			MediaLaunch:    adminUser.MediaLaunch,
			CreateLevel:    adminUser.CreateLevel,
			UsedStatus:     adminUser.UsedStatus,
			MainUserId:     adminUser.MainUserId,
			ContractType:   adminUser.ContractType,
			ProductVersion: 2,
		},
	}
	err = auth_token.CreateJwtToken(tokenBuilder, vars.Config.Auth.WechatMiniProgram, s.DbConnect)
	if err != nil {
		return nil, "", err
	}
	token = tokenBuilder.GetToken()
	return
}
