package user

import (
	"github.com/gin-gonic/gin"
	apiData "xiaoniuds.com/cid/api/data"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/user"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Api struct {
	C         *config.Config
	DbConnect *data.Data
}

func (a *Api) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData apiData.LoginData
		err := validator.BindJsonData(ctx, &loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		loginResponse := (&user.Service{C: a.C, DbConnect: a.DbConnect}).Login(loginData)
		if loginResponse != nil {
			response.Error(ctx, loginResponse)
			return
		}
		response.Success(ctx, loginResponse)
	}
}

func (a *Api) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	}
}
