package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/config"
)

type Server struct {
	c      *config.Config
	engine *gin.Engine
}

type Opt func(srv *Server)

func NewServer(c *config.Config) (srv *Server) {
	srv = &Server{
		c:      c,
		engine: gin.Default(),
	}

	srv.loadServes(
		NewUserServer(c),
	)

	return
}

func (srv *Server) loadServes(serves ...Opt) {
	for _, serve := range serves {
		serve(srv)
	}
}

func (srv *Server) Run() (err error) {
	err = srv.engine.Run(fmt.Sprintf(":%d", srv.c.Port))
	return
}
