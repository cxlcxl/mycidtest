package data

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"time"
	"xiaoniuds.com/cid/pkg/cache"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type DorisModel struct {
	db *gorm.DB
}

func NewDorisModel(connect string, connects *Data) *DorisModel {
	if connect == "" {
		connect = vars.DRDorisCid
	}
	return &DorisModel{
		db: connects.DbConnects[connect],
	}
}

func (m *DorisModel) QuerySQL(sql string, value interface{}, parameters ...interface{}) (err error) {
	err = m.db.Raw(sql, parameters...).Scan(value).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}
	}
	return
}

func (m *DorisModel) QuerySQLWithCache(
	cache cache.Cache, keySuffix, sql string, value interface{}, exp time.Duration, parameters ...interface{},
) (err *errs.MyErr) {
	v := cache.Get(keySuffix)
	if v != "" {
		e := json.Unmarshal([]byte(v), value)
		if e != nil {
			err = errs.Err(errs.SysError, e)
		}
	} else {
		e := m.QuerySQL(sql, value, parameters...)
		if e != nil {
			err = errs.Err(errs.SysError, e)
		} else {
			_ = cache.Set(keySuffix, value, exp)
		}
	}
	return
}

// QueryCallWithCache 带缓存的查询
func (m *DorisModel) QueryCallWithCache(
	cache cache.Cache, keySuffix string, value interface{}, exp time.Duration, fn QueryCallFunc, parameters ...interface{},
) (err *errs.MyErr) {
	v := cache.Get(keySuffix)
	if v != "" {
		e := json.Unmarshal([]byte(v), value)
		if e != nil {
			err = errs.Err(errs.SysError, e)
		}
	} else {
		err = fn(value, parameters...)
		if err == nil {
			_ = cache.Set(keySuffix, value, exp)
		}
	}
	return
}
