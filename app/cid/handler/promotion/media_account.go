package promotion

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
)

type MediaAccount struct {
	C         *config.Config
	DbConnect *data.Data
}

func (h *MediaAccount) TtActList() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *MediaAccount) KsActList() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}

func (h *MediaAccount) GdtActList() gin.HandlerFunc {
	return func(ctx *gin.Context) {}
}
