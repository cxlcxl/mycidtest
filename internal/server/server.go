package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/middleware"
	"xiaoniuds.com/cid/vars"
)

type Server struct {
	engine     *gin.Engine
	DbConnects *data.Data
}

type Opt func(srv *Server)

func NewServer() (srv *Server) {
	srv = &Server{
		DbConnects: data.NewDB(),
		engine:     gin.Default(),
	}

	srv.engine.Use(
		middleware.Cors(),
		middleware.RequestId(),
	)

	srv.loadServes(
		NewUserServer(),
		NewToolServer(),
		NewPromotionServer(),
		NewHomeReportServer(),
		// 对外 OpenApi
		NewOpenApiServer(),
		// 微信相关路由
		NewWechatServer(),
	)

	return
}

func (srv *Server) loadServes(serves ...Opt) {
	for _, serve := range serves {
		serve(srv)
	}
}

func (srv *Server) Run() (err error) {
	err = srv.engine.Run(fmt.Sprintf(":%d", vars.Config.Port))
	return
}
