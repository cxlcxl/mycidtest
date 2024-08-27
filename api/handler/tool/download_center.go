package tool

import (
	"github.com/gin-gonic/gin"
	apiData "xiaoniuds.com/cid/api/data"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/tool"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Tool struct {
	C         *config.Config
	DbConnect *data.Data
}

func (t *Tool) DownloadCenter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params apiData.DownloadCenterList
		if err := validator.BindJsonData(ctx, &params); err != nil {
			response.Error(ctx, err)
			return
		}
		logs, total, err := (&tool.Tool{
			C:         t.C,
			DbConnect: t.DbConnect,
		}).DownloadCenterList(params)

		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.PageSuccess(ctx, params.Page, params.PageSize, total, logs)
	}
}
