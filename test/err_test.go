package test

import (
	"errors"
	"testing"
	"xiaoniuds.com/cid/pkg/errs"
)

func TestErr(t *testing.T) {
	e := errs.Err(nil, errors.New("测试错误"), errors.New("6666666"))
	t.Log(e)
}
