package user

import (
	"fmt"
	apiData "xiaoniuds.com/cid/api/data"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
)

type Service struct {
	C         *config.Config
	DbConnect *data.Data
}

func (s *Service) Login(loginData apiData.LoginData) (err *errs.MyErr) {
	user, err := data.NewUserModel("", s.DbConnect).
		FindUserByLogin(loginData.Email, util.Password(loginData.Password, true))
	if err != nil {
		return
	}
	if user.IsLock != 0 {
		return errs.Err(errs.LoginUserExpireError)
	}

	// 判断是否过期
	if user.ParentId == 0 {

	}
	fmt.Println(user)
	return
}
