package auth_token

import (
	"github.com/golang-jwt/jwt/v5"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
)

func CreateLoginToken(user *data.User) (token *LoginToken, err *errs.MyErr) {
	claims := jwt.NewWithClaims(jwt.SigningMethodES256, jwt.MapClaims{
		"exp":            0,
		"iat":            0,
		"nbf":            0,
		"aud":            0,
		"user_id":        0,
		"user_name":      0,
		"user_full_name": 0,
		"email":          0,
		"mobile":         0,
		"project_id":     0,
		"group_id":       0,
		"used_status":    0,
		"create_level":   0,
	})

	signedString, e := claims.SignedString("")
	if e != nil {
		return nil, errs.Err(errs.ErrJwtSign, e)
	}
	token = &LoginToken{
		AccessToken: signedString,
		ExpireTime:  7200,
		UserInfo:    &LoginData{},
	}
	return
}

func ParseToken(token string) (user *data.User, err *errs.MyErr) {
	return
}
