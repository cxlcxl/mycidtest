package middleware

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/vars"
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
		builder := &auth_token.WebToken{
			Token: &auth_token.TokenInfo{
				AccessToken: header.Authorization,
			},
		}
		if err := auth_token.ParseToken(builder, auth); err != nil {
			response.Error(ctx, err)
			return
		} else {
			ctx.Set(vars.LoginKey, builder.User)
		}
		ctx.Next()
	}
}

func OpenApiAuth(auth config.Auth) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var header AuthHeader
		if err := ctx.ShouldBindHeader(&header); err != nil {
			response.Error(ctx, errs.Err(errs.ErrAuthFail, err))
			return
		}
		builder := &auth_token.OpenApiToken{
			Token: &auth_token.TokenInfo{
				AccessToken: header.Authorization,
			},
		}
		if err := auth_token.ParseToken(builder, auth); err != nil {
			response.Error(ctx, err)
			return
		} else {
			ctx.Set(vars.OpenApiLoginKey, builder.Data)
		}
		ctx.Next()
	}
}
