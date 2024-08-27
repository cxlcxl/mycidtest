package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/errs"
)

func CreateLoginToken(user *LoginData, auth config.Auth) (token *LoginToken, err *errs.MyErr) {
	exp := time.Hour * time.Duration(auth.Exp)
	loginClaims := &LoginClaims{
		UserInfo: user,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "xiaoniuds.com",
			Subject:   "cid",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			ID:        auth.Id,
			//Audience:  jwt.ClaimStrings{"cid"}, // 接收者
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, loginClaims)

	signedString, e := claims.SignedString([]byte(auth.SignKey))
	if e != nil {
		return nil, errs.Err(errs.ErrJwtToken, e)
	}

	token = &LoginToken{
		Token: TokenInfo{
			AccessToken: signedString,
			ExpireTime:  auth.Exp * 3600,
		},
		UserInfo: loginClaims.UserInfo,
	}
	return
}

func ParseToken(token string, auth config.Auth) (user *LoginData, err *errs.MyErr) {
	var loginClaims LoginClaims
	_, e := jwt.ParseWithClaims(token, &loginClaims, func(token *jwt.Token) (interface{}, error) {
		return []byte(auth.SignKey), nil
	})
	if e != nil {
		return nil, errs.Err(errs.ErrParseJwtToken, e)
	}

	// 数据库或缓存验证
	if err = dbCheckToken(token, loginClaims.UserInfo); err != nil {
		return nil, err
	}

	user = loginClaims.UserInfo
	return
}

func dbCheckToken(token string, user *LoginData) (err *errs.MyErr) {
	return nil
}
