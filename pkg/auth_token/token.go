package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

type TokenBuilder interface {
	TokenType() string
	BuildLoginClaims(jwt.RegisteredClaims) jwt.Claims
	SetTokenRsp(TokenInfo, *data.Data)
	GetToken() string
	MakeLoginClaims() jwt.Claims
	DbCheckToken(interface{}, *data.Data) *errs.MyErr
}

func CreateJwtToken(builder TokenBuilder, auth config.Auth, connects *data.Data) (err *errs.MyErr) {
	exp := time.Hour * time.Duration(auth.Exp)
	expireAt := time.Now().Add(exp)
	jwtClaims := jwt.RegisteredClaims{
		Issuer:    "xiaoniuds.com",
		Subject:   "cid",
		ExpiresAt: jwt.NewNumericDate(expireAt),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		ID:        auth.Id,
		//Audience:  jwt.ClaimStrings{"cid"}, // 接收者
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, builder.BuildLoginClaims(jwtClaims))

	signedString, e := claims.SignedString([]byte(auth.SignKey))
	if e != nil {
		return errs.Err(errs.ErrJwtToken, e)
	}

	builder.SetTokenRsp(TokenInfo{
		AccessToken: signedString,
		ExpireTime:  auth.Exp * 3600,
	}, connects)
	return
}

func ParseToken(builder TokenBuilder, auth config.Auth, connects *data.Data) (err *errs.MyErr) {
	if builder.GetToken() == "" {
		return errs.Err(errs.ParamError, errs.ErrMissToken)
	}
	loginClaims := builder.MakeLoginClaims()
	_, e := jwt.ParseWithClaims(builder.GetToken(), loginClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.SignKey), nil
	})
	if e != nil {
		return errs.Err(errs.ErrParseJwtToken, e)
	}

	// 数据库或缓存验证
	if err = builder.DbCheckToken(loginClaims, connects); err != nil {
		return err
	}

	return
}
