package handle

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/wechat/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/wechat"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Auth struct {
	DbConnect *data.Data
}

func (a *Auth) MiniProgramLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData statement.LoginData
		err := validator.BindJsonData(ctx, &loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		bindUsers, openId, err := (&wechat.Service{DbConnect: a.DbConnect}).GetUsersByCode(loginData.Code)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, gin.H{"users": bindUsers, "openid": openId})
	}
}

func (a *Auth) BindUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var bindData statement.BindUser
		err := validator.BindJsonData(ctx, &bindData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		bindUsers, err := (&wechat.Service{DbConnect: a.DbConnect}).BindUser(bindData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, gin.H{"users": bindUsers, "openid": bindData.OpenId})
	}
}

func (a *Auth) SelectUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var selectData statement.SelectUser
		err := validator.BindJsonData(ctx, &selectData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		user, token, err := (&wechat.Service{DbConnect: a.DbConnect}).SelectUserLogin(selectData.OpenId, selectData.UserId)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, gin.H{"user": user, "access_token": token})
	}
}
