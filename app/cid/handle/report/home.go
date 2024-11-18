package report

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/cid/report"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Home struct {
	C         *config.Config
	DbConnect *data.Data
}

func (h *Home) OrderSum() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.ReportHomeOrderSum
		if err := validator.BindJsonData(ctx, &params); err != nil {
			response.Error(ctx, err)
			return
		}
		orderSum, err := (&report.HomeService{
			C:         h.C,
			DbConnect: h.DbConnect,
		}).OrderSum(params)
		if err != nil {
			response.Error(ctx, err)
			return
		}
		response.Success(ctx, orderSum)
	}
}
