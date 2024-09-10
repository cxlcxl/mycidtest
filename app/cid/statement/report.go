package statement

import "time"

type ReportHomeOrderSum struct {
	StartDate time.Time `form:"start_date" binding:"required" time_format:"2006-01-02"`
	EndDate   time.Time `form:"end_date" binding:"required" time_format:"2006-01-02"`
	MediaType int       `form:"media_type" binding:"numeric"`
	Platform  int       `form:"platform" binding:"numeric"`
}
