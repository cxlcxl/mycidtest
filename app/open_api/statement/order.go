package statement

import (
	"time"
	"xiaoniuds.com/cid/app/cid/statement"
)

type OrderList struct {
	StartTime time.Time `form:"start_time"`
	EndTime   time.Time `form:"end_time"`
	ShopType  int       `form:"shop_type"`
	TimeType  int       `form:"time_type"`
	IsHidden  int       `form:"is_hidden"`
	*statement.Pagination
}
