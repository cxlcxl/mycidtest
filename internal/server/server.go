package server

import (
	"fmt"
	"xiaoniuds.com/cid/internal/middleware"
	"xiaoniuds.com/cid/pkg/mylog"

	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
)

type Server struct {
	C          *config.Config
	engine     *gin.Engine
	DbConnects *data.Data
	Log        *mylog.Log
}

type Opt func(srv *Server)

func NewServer(c *config.Config) (srv *Server) {
	myLog := mylog.NewLog()
	srv = &Server{
		C:          c,
		DbConnects: data.NewDB(c, myLog),
		engine:     gin.Default(),
		Log:        myLog,
	}

	srv.engine.Use(
		middleware.Cors(),
		middleware.RequestId(),
	)

	srv.loadServes(
		NewOpenApiServer(), // 对外 OpenApi
		NewUserServer(),
		NewToolServer(),
		NewPromotionServer(),
		NewHomeReportServer(),
	)

	return
}

func (srv *Server) loadServes(serves ...Opt) {
	for _, serve := range serves {
		serve(srv)
	}
}

func (srv *Server) Run() (err error) {
	err = srv.engine.Run(fmt.Sprintf(":%d", srv.C.Port))
	return
}
