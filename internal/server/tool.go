package server

import (
	"xiaoniuds.com/cid/app/cid/handle/tool"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewToolServer() Opt {
	return func(srv *Server) {
		toolApi := &tool.Tool{DbConnect: srv.DbConnects}

		group := srv.engine.Group("/tools", middleware.LoginAuth(srv.DbConnects))
		{
			group.GET("/download_center", toolApi.DownloadCenter())
		}
	}
}
