package auth_token

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
)

type WebToken struct {
	User  *LoginData `json:"user"`
	Token *TokenInfo `json:"token"`
}

func (t *WebToken) TokenType() string {
	return "web"
}

func (t *WebToken) BuildLoginClaims(claims jwt.RegisteredClaims) jwt.Claims {
	loginClaims := &LoginClaims{
		UserInfo:         t.User,
		RegisteredClaims: claims,
	}
	return loginClaims
}

func (t *WebToken) SetTokenRsp(tokenInfo TokenInfo, connects *data.Data) {
	t.Token = &tokenInfo
	_ = base.NewACTokenModel("", connects).Save(&base.ACToken{
		MainUserId:     t.User.MainUserId,
		UserId:         t.User.MainUserId,
		IP:             t.User.IP,
		Scopes:         "*",
		TokenType:      t.TokenType(),
		AccessToken:    tokenInfo.AccessToken,
		AccessTokenMD5: util.Md5(tokenInfo.AccessToken),
		RefreshToken:   "",
		ExpireTime:     time.Now().Add(time.Second * time.Duration(tokenInfo.ExpireTime)),
		CreateTime:     time.Now(),
	})
	return
}

func (t *WebToken) GetToken() string {
	return t.Token.AccessToken
}

func (t *WebToken) MakeLoginClaims() jwt.Claims {
	return &LoginClaims{}
}

func (t *WebToken) DbCheckToken(claims interface{}, connects *data.Data) *errs.MyErr {
	t.User = claims.(*LoginClaims).UserInfo
	token, err := base.NewACTokenModel("", connects).GetOneByBuilder(func(query *gorm.DB) *gorm.DB {
		return query.
			Where("access_token_md5 = ?", util.Md5(t.Token.AccessToken))
	}, []string{"id", "expire_time"})
	if err != nil {
		return err
	}
	if token == nil {
		return errs.Err(errs.ErrParseJwtToken, errors.New("token is invalid"))
	}
	if token.ExpireTime.Before(time.Now()) {
		return errs.Err(errs.ErrParseJwtToken, errors.New("token is expired"))
	}
	return nil
}
