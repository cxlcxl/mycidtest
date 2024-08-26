package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaoniuds.com/cid/pkg/errs"
)

func Success(ctx *gin.Context, data interface{}) {
	ctx.JSON(http.StatusOK, gin.H{"code": 0, "msg": "ok", "data": data})
}

func Error(ctx *gin.Context, err *errs.MyErr) {
	ctx.JSON(http.StatusOK, gin.H{"code": err.Code(), "msg": err.Error(), "data": nil})
}
