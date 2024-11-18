package tool

import (
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/task"
	"xiaoniuds.com/cid/pkg/errs"
)

type Tool struct {
	C         *config.Config
	DbConnect *data.Data
}

func (t *Tool) DownloadCenterList(params statement.DownloadCenterList) (
	logs []*task.DownloadCenterListItem, total int64, err *errs.MyErr,
) {
	logs, total, err = task.NewDownloadCenterModel("", t.DbConnect).
		GetDownloadCenterList(params.LoginData, params.TaskName, params.TaskType, params.Page, params.PageSize)
	return
}
