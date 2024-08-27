package data

import "xiaoniuds.com/cid/pkg/auth_token"

type DownloadCenterList struct {
	*Pagination
	LoginData *auth_token.LoginData

	TaskName string `form:"task_name"`
	TaskType int    `form:"task_type"`
}
