package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

func RequestId() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		requestId := fmt.Sprintf("%d-%s", time.Now().UnixNano(), util.RandString(20))
		ctx.Set(vars.RequestIdKey, util.Sha1(requestId))
		ctx.Next()
	}
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
