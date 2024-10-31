package server

import (
	"xiaoniuds.com/cid/app/cid/handler/report"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewHomeReportServer() Opt {
	return func(srv *Server) {
		homeReportApi := &report.Home{C: srv.C, DbConnect: srv.DbConnects}

		r := srv.engine.Group("/report", middleware.LoginAuth(srv.C.Auth.Login, srv.DbConnects))
		{
			r.GET("/order_sum", homeReportApi.OrderSum())
		}
	}
}
