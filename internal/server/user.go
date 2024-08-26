package server

import (
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/api/user"
)

func NewUserServer(c *config.Config) Opt {
	userApi := &user.Api{C: c}

	return func(srv *Server) {
		group := srv.engine.Group("/user")
		{
			group.POST("/list", userApi.Login())
			group.POST("/", userApi.Create())
		}
	}
}
