package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
)

type Service struct {
	C         *config.Config
	DbConnect *data.Data
}

func (s *Service) GetToken(params statement.Token) (builder *auth_token.OpenApiToken, err *errs.MyErr) {
	var app *config.OpenApiApp
	for _, a := range s.C.Auth.OpenApiApps {
		if a.AppId == params.AppId {
			if a.AppSecret != params.AppSecret {
				err = errs.Err(errs.SysError, errs.ErrJwtToken)
				return
			}
			app = &a
		}
	}
	if app == nil {
		err = errs.Err(errs.SysError, errs.ErrMissUserInfo)
		return
	}
	builder = &auth_token.OpenApiToken{
		Data: &auth_token.OpenApiData{
			MainUserId: app.MainUserId,
			AppId:      params.AppId,
			AppSecret:  params.AppSecret,
		},
	}
	err = auth_token.CreateJwtToken(builder, s.C.Auth.OpenApi)

	return
}
