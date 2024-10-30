package statement

import (
	"time"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/pkg/auth_token"
)

type OrderList struct {
	StartTime time.Time `form:"start_time"`
	EndTime   time.Time `form:"end_time"`
	ShopType  int       `form:"shop_type"`
	TimeType  int       `form:"time_type"`
	IsHidden  int       `form:"is_hidden"`
	*statement.Pagination
	OpenApiLoginData *auth_token.OpenApiLoginData
}
