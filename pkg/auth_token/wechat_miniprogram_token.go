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

type WechatMiniProgramToken struct {
	Data  *LoginData `json:"data"`
	Token *TokenInfo `json:"token"`
}

func (t *WechatMiniProgramToken) TokenType() string {
	return "wechat_mini_program"
}

func (t *WechatMiniProgramToken) BuildLoginClaims(claims jwt.RegisteredClaims) jwt.Claims {
	loginClaims := &WechatMiniProgramClaims{
		WechatMiniProgramData: t.Data,
		RegisteredClaims:      claims,
	}
	return loginClaims
}

func (t *WechatMiniProgramToken) SetTokenRsp(tokenInfo TokenInfo, connects *data.Data) {
	t.Token = &tokenInfo
	_ = base.NewACTokenModel("", connects).Save(&base.ACToken{
		MainUserId:     t.Data.MainUserId,
		UserId:         t.Data.MainUserId,
		IP:             t.Data.IP,
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

func (t *WechatMiniProgramToken) GetToken() string {
	return t.Token.AccessToken
}

func (t *WechatMiniProgramToken) MakeLoginClaims() jwt.Claims {
	return &OpenApiClaims{}
}

func (t *WechatMiniProgramToken) DbCheckToken(claims interface{}, connects *data.Data) *errs.MyErr {
	t.Data = claims.(*WechatMiniProgramClaims).WechatMiniProgramData
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
