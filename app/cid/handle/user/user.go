package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/cid/user"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
	"xiaoniuds.com/cid/vars"
)

type Api struct {
	DbConnect *data.Data
}

func (a *Api) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData statement.LoginData
		err := validator.BindJsonData(ctx, &loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		loginInfo, err := (&user.Service{DbConnect: a.DbConnect}).Login(loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, loginInfo)
	}
}

func (a *Api) ZoneDomain() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.ZoneDomain
		err := validator.BindJsonData(ctx, &params)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		zone, err := (&user.Service{DbConnect: a.DbConnect}).ZoneDomain(params)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		if zone == 0 {
			response.Success(ctx, fmt.Sprintf("cli.%s", vars.Config.MainDomain))
		} else {
			response.Success(ctx, fmt.Sprintf("cli%d.%s", zone, vars.Config.MainDomain))
		}
	}
}

func (a *Api) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	}
}
