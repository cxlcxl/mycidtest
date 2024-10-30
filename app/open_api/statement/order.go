package statement

import (
	"time"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/pkg/auth_token"
)

type OrderList struct {
	StartTime time.Time `form:"start_time" binding:"required"`
	EndTime   time.Time `form:"end_time" binding:"required"`
	ShopType  int       `form:"shop_type" binding:"required"`
	TimeType  int       `form:"time_type" binding:"required"`
	IsHidden  int       `form:"is_hidden"`
	*statement.Pagination
	OpenApiData *auth_token.OpenApiData
}
