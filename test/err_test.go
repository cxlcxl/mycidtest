package test

import (
	"errors"
	"fmt"
	"testing"
	"time"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
)

func TestErr(t *testing.T) {
	e := errs.Err(nil, errors.New("测试错误"), errors.New("6666666"))
	t.Log(e)
}

func TestUtil(t *testing.T) {
	requestId := fmt.Sprintf("%d-%s", time.Now().UnixNano(), util.RandString(20))
	t.Log(requestId)
}

func TestDiyTime(t *testing.T) {
	parse, _ := time.Parse(time.DateTime, "2024-10-01 10:00:00")
	dateTime := data.DbDateTime(parse)
	value, _ := dateTime.Value()
	if value != nil {
		fmt.Println(value.(time.Time).Format(time.DateTime))
	}

}
