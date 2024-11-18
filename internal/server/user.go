package server

import (
	"xiaoniuds.com/cid/app/cid/handle/user"
)

func NewUserServer() Opt {
	return func(srv *Server) {
		userApi := &user.Api{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/user")
		{
			group.GET("/zone_domain", userApi.ZoneDomain())
			group.POST("/login", userApi.Login())
			group.POST("/", userApi.Create())
		}
	}
}
