package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type Service struct {
	DbConnect *data.Data
}

func (s *Service) GetToken(params statement.Token) (builder *auth_token.OpenApiToken, err *errs.MyErr) {
	var app *config.OpenApiApp
	for _, a := range vars.Config.Auth.OpenApiApps {
		if a.AppId == params.AppId {
			if a.AppSecret != params.AppSecret {
				err = errs.Err(errs.SysError, errs.OpenApiErrWornSecret)
				return
			}
			app = &a
		}
	}
	if app == nil {
		err = errs.Err(errs.SysError, errs.OpenApiErrMissAppId)
		return
	}
	builder = &auth_token.OpenApiToken{
		Data: &auth_token.OpenApiData{
			MainUserId: app.MainUserId,
			AppId:      params.AppId,
		},
	}
	err = auth_token.CreateJwtToken(builder, vars.Config.Auth.OpenApi, s.DbConnect)

	return
}
