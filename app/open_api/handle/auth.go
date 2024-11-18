package handle

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/open_api"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Auth struct {
	DbConnect *data.Data
}

func (a *Auth) GetToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData statement.Token
		err := validator.BindJsonData(ctx, &loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		loginInfo, err := (&open_api.Service{DbConnect: a.DbConnect}).GetToken(loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, loginInfo)
	}
}
