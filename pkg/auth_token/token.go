package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"xiaoniuds.com/cid/pkg/errs"
)

type WebToken struct {
	User  *LoginData
	Token *TokenInfo
}

func (t *WebToken) BuildLoginClaims(claims jwt.RegisteredClaims) jwt.Claims {
	loginClaims := &LoginClaims{
		UserInfo:         t.User,
		RegisteredClaims: claims,
	}
	return loginClaims
}

func (t *WebToken) SetTokenRsp(tokenInfo TokenInfo) {
	t.Token = &tokenInfo
	return
}

func (t *WebToken) GetToken() string {
	return t.Token.AccessToken
}

func (t *WebToken) MakeLoginClaims() jwt.Claims {
	return &LoginClaims{}
}

func (t *WebToken) DbCheckToken(claims interface{}) *errs.MyErr {
	t.User = claims.(*LoginClaims).UserInfo
	return nil
}

type OpenApiToken struct {
	Data  *OpenApiData
	Token *TokenInfo
}

func (t *OpenApiToken) BuildLoginClaims(claims jwt.RegisteredClaims) jwt.Claims {
	loginClaims := &OpenApiClaims{
		OpenApiData:      t.Data,
		RegisteredClaims: claims,
	}
	return loginClaims
}

func (t *OpenApiToken) SetTokenRsp(tokenInfo TokenInfo) {
	t.Token = &tokenInfo
	return
}

func (t *OpenApiToken) GetToken() string {
	return t.Token.AccessToken
}

func (t *OpenApiToken) MakeLoginClaims() jwt.Claims {
	return &OpenApiClaims{}
}

func (t *OpenApiToken) DbCheckToken(claims interface{}) *errs.MyErr {
	t.Data = claims.(*OpenApiClaims).OpenApiData
	return nil
}
