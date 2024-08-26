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
	SysError             = &MyErr{400, "系统繁忙"}
	ParamError           = &MyErr{400, "参数解析错误"}
	ConfigLoadError      = &MyErr{400, "配置加载错误"}
	LoginFinUserError    = &MyErr{401, "用户名或密码错误"}
	LoginUserExpireError = &MyErr{400, "该账户已过期，请联系运营同学"}
	ErrJwtSign           = &MyErr{500, "TOKEN 生成失败"}
)

func (ce *MyErr) Error() string {
	return ce.message
}

func (ce *MyErr) Code() int {
	return ce.errCode
}

func (ce *MyErr) join(errs ...error) string {
	if ce.message != "" {
		ce.message = ce.message + ": "
	}

	var errMsg []string
	for _, err := range errs {
		if err != nil {
			errMsg = append(errMsg, err.Error())
		}
	}

	return fmt.Sprintf("%s%s", ce.message, strings.Join(errMsg, "; "))
}

func Err(myErr *MyErr, errs ...error) (err *MyErr) {
	if myErr != nil {
		if len(errs) == 0 {
			return nil
		}
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
