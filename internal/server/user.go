package server

import (
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/middleware"
	"xiaoniuds.com/cid/internal/service/user"
)

func NewUserServer(c *config.Config) Opt {
	userService := &user.Service{C: c}

	return func(srv *Server) {
		group := srv.engine.Group("/user")
		{
			group.Use(middleware.LoginFailLimit()).POST("/list", userService.Login())
			group.POST("/", userService.Create())
		}
	}
}
