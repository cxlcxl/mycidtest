package promotion

import (
	"errors"
	"github.com/gin-gonic/gin"
	"slices"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/cid/promotion"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Promotion struct {
	DbConnect *data.Data
}

func (h *Promotion) PddGoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.PddGoodsList
		if err := validator.BindJsonData(ctx, &params); err != nil {
			response.Error(ctx, err)
			return
		}

		if params.SortField != "" && !slices.Contains([]string{"create_time", "sale_num", "today_sale_num"}, params.SortField) {
			response.Error(ctx, errs.Err(errs.ParamError, errors.New("不支持的排序字段")))
			return
		}

		list, total, err := (&promotion.PddGoods{DbConnect: h.DbConnect}).List(params)
		if err != nil {
			response.Error(ctx, err)
			return
		}
		// 是否有备案推广位,没有则进入推广位授权
		//PddPidService::getValidPidByUserInfo($userInfo);
		response.PageSuccess(ctx, params.Page, params.PageSize, total, list)
	}
}

func (h *Promotion) TbGoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *Promotion) JdGoods() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
