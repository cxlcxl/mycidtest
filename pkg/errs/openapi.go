package errs

var (
	OpenApiErrMissAppId    = &MyErr{2001, "AppId 不存在"}
	OpenApiErrWornSecret   = &MyErr{2001, "AppSecret 错误"}
	OpenApiErrUserNotMatch = &MyErr{2001, "用户信息不匹配"}
)
