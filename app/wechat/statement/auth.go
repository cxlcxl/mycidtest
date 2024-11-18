package statement

type LoginData struct {
	Code string `json:"code" binding:"required"`
}

type BindUser struct {
	OpenId string `json:"openid" binding:"required,alphanum"`
	Email  string `json:"email" binding:"required,email"`
	Pass   string `json:"pass" binding:"required,password"`
}

type SelectUser struct {
	OpenId string `json:"openid" binding:"required,alphanum"`
	UserId int64  `json:"user_id" binding:"required,numeric"`
}
