package user

import (
	"fmt"
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

		loginInfo, err := (&user.Service{C: a.C, DbConnect: a.DbConnect}).Login(loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, loginInfo)
	}
}

func (a *Api) ZoneDomain() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params apiData.ZoneDomain
		err := validator.BindJsonData(ctx, &params)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		zone, err := (&user.Service{C: a.C, DbConnect: a.DbConnect}).ZoneDomain(params)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		if zone == 0 {
			response.Success(ctx, fmt.Sprintf("cli.%s", a.C.MainDomain))
		} else {
			response.Success(ctx, fmt.Sprintf("cli%d.%s", zone, a.C.MainDomain))
		}
	}
}

func (a *Api) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	}
}
