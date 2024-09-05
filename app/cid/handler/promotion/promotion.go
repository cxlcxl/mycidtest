package promotion

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Promotion struct {
	C         *config.Config
	DbConnect *data.Data
}

func (h *Promotion) PddGoodsLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.PddGoodsList
		if err := validator.BindJsonData(ctx, &params); err != nil {
			response.Error(ctx, err)
			return
		}

		response.PageSuccess(ctx, params.Page, params.PageSize, 0, nil)
	}
}

func (h *Promotion) TbGoodsLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *Promotion) JdGoodsLink() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
