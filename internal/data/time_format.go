package data

import (
	"database/sql/driver"
	"fmt"
	"time"
)

// DbDateTime 数据库时间格式化
type DbDateTime time.Time
type DbDate time.Time

type Timestamp struct {
	CreateTime DbDateTime `json:"create_time" gorm:"column:create_time"`
	UpdateTime DbDateTime `json:"update_time" gorm:"column:update_time"`
}

func (t *DbDateTime) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateTime))), nil
}

func (t DbDateTime) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *DbDateTime) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DbDateTime(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
func (t *DbDate) MarshalJSON() ([]byte, error) {
	tTime := time.Time(*t)
	return []byte(fmt.Sprintf("\"%v\"", tTime.Format(time.DateOnly))), nil
}

func (t DbDate) Value() (driver.Value, error) {
	var zeroTime time.Time
	tlt := time.Time(t)
	//判断给定时间是否和默认零时间的时间戳相同
	if tlt.UnixNano() == zeroTime.UnixNano() {
		return nil, nil
	}
	return tlt, nil
}

func (t *DbDate) Scan(v interface{}) error {
	if value, ok := v.(time.Time); ok {
		*t = DbDate(value)
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}
