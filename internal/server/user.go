package server

import (
	"xiaoniuds.com/cid/api/handler/user"
)

func NewUserServer() Opt {
	return func(srv *Server) {
		userApi := &user.Api{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/user")
		{
			group.POST("/login", userApi.Login())
			group.POST("/", userApi.Create())
		}
	}
}
