package middleware

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/constant"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util/response"
)

func LoginFailLimit() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}

type AuthHeader struct {
	Authorization string `header:"Authorization"`
}

func LoginAuth(auth config.Auth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header AuthHeader
		if err := ctx.ShouldBindHeader(&header); err != nil {
			response.Error(ctx, errs.Err(errs.ErrAuthFail, err))
			return
		}
		if loginData, err := auth_token.ParseToken(header.Authorization, auth); err != nil {
			response.Error(ctx, err)
			return
		} else {
			ctx.Set(constant.LoginKey, loginData)
		}
		ctx.Next()
	}
}
