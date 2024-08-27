package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

func CreateLoginToken(user *data.User, auth config.Auth) (token *LoginToken, err *errs.MyErr) {
	exp := time.Hour * time.Duration(auth.Exp)
	loginClaims := &LoginClaims{
		UserInfo: &LoginData{
			UserId:       user.UserId,
			ProjectId:    user.ProjectId,
			GroupId:      user.GroupId,
			Email:        user.Email,
			UserName:     user.UserName,
			UserFullName: user.UserFullName,
			Mobile:       user.Mobile,
			DataRange:    user.DataRange,
			LatestNews:   user.LatestNews,
			CompanyType:  user.CompanyType,
			Industry:     user.Industry,
			MediaLaunch:  user.MediaLaunch,
			CreateLevel:  user.CreateLevel,
			UsedStatus:   user.UsedStatus,
			MainUserId:   user.MainUserId,
			ContractType: user.ContractType,
		},
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

	user = loginClaims.UserInfo
	return
}
