package errs

import (
	"errors"
	"net/http"
)

type MyErr struct {
	errCode int
	message string
}

var (
	SysError        = &MyErr{400, "系统繁忙"}
	ParamError      = &MyErr{400, "参数解析错误"}
	ConfigLoadError = &MyErr{400, "配置加载错误"}
)

func (ce *MyErr) Error() string {
	return ce.message
}

func (ce *MyErr) Code() int {
	return ce.errCode
}

func joinErrsMsg(prefix string, errs ...error) string {
	if prefix != "" {
		prefix = prefix + ": "
	}
	return prefix + errors.Join(errs...).Error()
}

func Err(myErr *MyErr, errs ...error) (err *MyErr) {
	if myErr != nil {
		if len(errs) == 0 {
			return nil
		}
		err = &MyErr{
			errCode: myErr.errCode,
			message: joinErrsMsg(myErr.message, errs...),
		}
	} else {
		if len(errs) > 0 {
			err = &MyErr{
				errCode: http.StatusBadRequest,
				message: joinErrsMsg("", errs...),
			}
		}
	}
	return
}
