package server

import (
	"xiaoniuds.com/cid/app/open_api/handle"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewOpenApiServer() Opt {
	return func(srv *Server) {
		api := &handle.OpenOrder{DbConnect: srv.DbConnects}
		auth := &handle.Auth{DbConnect: srv.DbConnects}

		group := srv.engine.Group("/open_api")
		{
			group.POST("/token", auth.GetToken())

			use := group.Use(middleware.OpenApiAuth(srv.DbConnects))
			{
				use.GET("/order/get", api.OrderList())
			}
		}
	}
}
