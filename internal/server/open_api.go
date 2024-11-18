package server

import (
	"xiaoniuds.com/cid/app/open_api/handle"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewOpenApiServer() Opt {
	return func(srv *Server) {
		api := &handle.OpenOrder{C: srv.C, DbConnect: srv.DbConnects}
		auth := &handle.Auth{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/open_api")
		{
			group.POST("/token", auth.GetToken())

			use := group.Use(middleware.OpenApiAuth(srv.C.Auth.OpenApi, srv.DbConnects))
			{
				use.GET("/order/get", api.OrderList())
			}
		}
	}
}
