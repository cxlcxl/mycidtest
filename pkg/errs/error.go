package errs

import (
	"fmt"
	"net/http"
	"strings"
)

type MyErr struct {
	errCode int
	message string
}

var (
	SysError             = &MyErr{1400, "系统繁忙"}
	ParamError           = &MyErr{1400, "参数解析错误"}
	ConfigLoadError      = &MyErr{1400, "配置加载错误"}
	LoginEmailNotExist   = &MyErr{1400, "登录邮箱不存在"}
	LoginFinUserError    = &MyErr{1401, "用户名或密码错误"}
	LoginUserExpireError = &MyErr{1400, "该账户已过期，请联系运营同学"}
	ErrMissToken         = &MyErr{1500, "TOKEN 缺失"}
	ErrJwtToken          = &MyErr{1500, "TOKEN 生成失败"}
	ErrParseJwtToken     = &MyErr{1500, "TOKEN 解析失败"}
	ErrAuthFail          = &MyErr{15001, "TOKEN 验证失败"}
	ErrMissUserInfo      = &MyErr{15001, "用户信息缺失"}
	ErrJsonUnmarshal     = &MyErr{15001, "JSON 解码失败"}
	ErrJsonMarshal       = &MyErr{15001, "JSON 加码失败"}
)

func (ce *MyErr) Error() string {
	return ce.message
}

func (ce *MyErr) Code() int {
	return ce.errCode
}

func (ce *MyErr) join(errs ...error) string {
	var errMsg []string
	for _, err := range errs {
		if err != nil {
			errMsg = append(errMsg, err.Error())
		}
	}

	return fmt.Sprintf("%s %s", ce.message, strings.Join(errMsg, "; "))
}

func Err(myErr *MyErr, errs ...error) (err *MyErr) {
	if myErr != nil {
		err = &MyErr{
			errCode: myErr.errCode,
			message: myErr.join(errs...),
		}
	} else {
		if len(errs) > 0 {
			err = &MyErr{
				errCode: http.StatusBadRequest,
				message: (&MyErr{0, ""}).join(errs...),
			}
		}
	}
	return
}
