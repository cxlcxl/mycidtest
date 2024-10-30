package server

import (
	"xiaoniuds.com/cid/app/open_api/handler"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewOpenApiServer() Opt {
	return func(srv *Server) {
		api := &handler.OpenOrder{C: srv.C, DbConnect: srv.DbConnects}
		auth := &handler.Auth{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/open_api")
		{
			group.POST("/token", auth.GetToken())

			use := group.Use(middleware.OpenApiAuth(srv.C.Auth.OpenApi))
			{
				use.GET("/order/get", api.OrderList())
			}
		}
	}
}
