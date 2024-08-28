package server

import (
	"xiaoniuds.com/cid/app/open_api/handler"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewOpenApiServer() Opt {
	return func(srv *Server) {
		api := &handler.OpenOrder{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/api", middleware.OpenApiAuth(srv.C.Auth.OpenApi))
		{
			group.GET("/order", api.OrderList())
		}
	}
}
