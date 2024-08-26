package user

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type Service struct {
	C *config.Config
}

type LoginData struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (s *Service) Login() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var loginData LoginData
		err := validator.BindJsonData(ctx, &loginData)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.Success(ctx, nil)
	}
}

func (s *Service) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{"msg": "ok"})
	}
}
