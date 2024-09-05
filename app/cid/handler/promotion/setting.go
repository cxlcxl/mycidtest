package promotion

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
)

type Setting struct {
	C         *config.Config
	DbConnect *data.Data
}

func (h *Setting) Callback() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *Setting) Notice() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
