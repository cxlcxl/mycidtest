package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type OpenOrder struct {
	C         *config.Config
	DbConnect *data.Data
}

func (o *OpenOrder) OrderList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.OrderList
		if err := validator.BindJsonData(ctx, &params, func(d interface{}) error {
			end := d.(*statement.OrderList).EndTime.AddDate(0, 0, -1)
			if d.(*statement.OrderList).StartTime.Before(end) {
				return errors.New("最大查询 1 天的数据")
			}
			return nil
		}); err != nil {

		}
	}
}
