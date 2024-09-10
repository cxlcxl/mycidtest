package cache

import (
	"time"
	"xiaoniuds.com/cid/pkg/errs"
)

type Cache interface {
	Get(string) string
	Set(keySuffix string, value interface{}, exp time.Duration) *errs.MyErr
}
