package tool

import (
	apiData "xiaoniuds.com/cid/api/data"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	data2 "xiaoniuds.com/cid/internal/data/task"
	"xiaoniuds.com/cid/pkg/errs"
)

type Tool struct {
	C         *config.Config
	DbConnect *data.Data
}

func (t *Tool) DownloadCenterList(params apiData.DownloadCenterList) (
	logs []*data2.DownloadCenterListItem, total int64, err *errs.MyErr,
) {
	logs, total, err = data2.NewDownloadCenterModel("", t.DbConnect).
		GetDownloadCenterList(params.LoginData, params.TaskName, params.TaskType, params.Page, params.PageSize)
	return
}
