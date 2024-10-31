package statement

import (
	"time"
	"xiaoniuds.com/cid/pkg/auth_token"
)

type OrderList struct {
	StartTime   time.Time `form:"start_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	EndTime     time.Time `form:"end_time" binding:"required" time_format:"2006-01-02 15:04:05"`
	ShopType    int       `form:"shop_type" binding:"required,numeric"`
	TimeType    int       `form:"time_type" binding:"required,numeric"`
	Page        int       `json:"page" form:"page" binding:"required,numeric"`
	PageSize    int       `json:"page_size" form:"page_size" binding:"required,numeric"`
	OpenApiData *auth_token.OpenApiData
}
