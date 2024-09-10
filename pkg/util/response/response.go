package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

func RequestId(ctx *gin.Context) string {
	requestId, _ := ctx.Get(vars.RequestIdKey)
	return requestId.(string)
}

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"request_id": RequestId(ctx),
		"code":       0,
		"msg":        "ok",
		"data":       data,
	})
}

func PageSuccess(ctx *gin.Context, page, pageSize int, total int64, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"request_id": RequestId(ctx),
		"code":       0,
		"msg":        "ok",
		"data":       data,
		"page_info": gin.H{
			"page":        page,
			"page_size":   pageSize,
			"total_count": total,
		},
	})
}

func Error(ctx *gin.Context, err *errs.MyErr) {
	ctx.JSON(http.StatusOK, gin.H{
		"request_id": RequestId(ctx),
		"code":       err.Code(),
		"msg":        err.Error(),
		"data":       nil,
	})
	ctx.Abort()
}
