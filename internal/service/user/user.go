package user

import (
	apiData "xiaoniuds.com/cid/api/data"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	data2 "xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
)

type Service struct {
	C         *config.Config
	DbConnect *data.Data
}

func (s *Service) ZoneDomain(params apiData.ZoneDomain) (zone int, err *errs.MyErr) {
	userZone, err := data2.NewUserZoneModel("", s.DbConnect).FindUserZoneByEmail(params.Email)
	if err != nil {
		return
	}
	if userZone == nil {
		return -1, errs.Err(errs.LoginEmailNotExist)
	}
	zone = userZone.ZoneIndex
	return
}

func (s *Service) Login(params apiData.LoginData) (loginInfo *auth_token.LoginToken, err *errs.MyErr) {
	user, err := data.NewUserModel("", s.DbConnect).
		FindUserByLogin(params.Email, util.Password(params.Password, false))
	if err != nil {
		return
	}
	if user.IsLock != 0 {
		return nil, errs.Err(errs.LoginUserExpireError)
	}

	// 判断是否过期
	if user.ParentId == 0 {

	}

	loginData := &auth_token.LoginData{
		UserId:         user.UserId,
		ProjectId:      user.ProjectId,
		GroupId:        user.GroupId,
		Email:          user.Email,
		UserName:       user.UserName,
		UserFullName:   user.UserFullName,
		Mobile:         user.Mobile,
		DataRange:      user.DataRange,
		LatestNews:     user.LatestNews,
		CompanyType:    user.CompanyType,
		Industry:       user.Industry,
		MediaLaunch:    user.MediaLaunch,
		CreateLevel:    user.CreateLevel,
		UsedStatus:     user.UsedStatus,
		MainUserId:     user.MainUserId,
		ContractType:   user.ContractType,
		ProductVersion: params.ProductVersion,
	}
	loginInfo, err = auth_token.CreateLoginToken(loginData, s.C.Auth.Login)
	return
}
