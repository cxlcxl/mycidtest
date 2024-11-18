package tool

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/cid/tool"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Tool struct {
	DbConnect *data.Data
}

func (t *Tool) DownloadCenter() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.DownloadCenterList
		if err := validator.BindJsonData(ctx, &params); err != nil {
			response.Error(ctx, err)
			return
		}
		logs, total, err := (&tool.Tool{
			DbConnect: t.DbConnect,
		}).DownloadCenterList(params)

		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.PageSuccess(ctx, params.Page, params.PageSize, total, logs)
	}
}
