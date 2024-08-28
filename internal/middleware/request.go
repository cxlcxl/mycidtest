package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"xiaoniuds.com/cid/constant"
	"xiaoniuds.com/cid/pkg/util"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := fmt.Sprintf("%d-%s", time.Now().UnixNano(), util.RandString(20))
		ctx.Set(constant.RequestIdKey, util.Sha1(requestId))
		ctx.Next()
	}
}
