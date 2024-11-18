package server

import (
	"xiaoniuds.com/cid/app/cid/handle/report"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewHomeReportServer() Opt {
	return func(srv *Server) {
		homeReportApi := &report.Home{DbConnect: srv.DbConnects}

		r := srv.engine.Group("/report", middleware.LoginAuth(srv.DbConnects))
		{
			r.GET("/order_sum", homeReportApi.OrderSum())
		}
	}
}
