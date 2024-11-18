package server

import (
	"xiaoniuds.com/cid/app/wechat/handle"
)

func NewWechatServer() Opt {
	return func(srv *Server) {
		auth := &handle.Auth{C: srv.C, DbConnect: srv.DbConnects}

		group := srv.engine.Group("/api/wx")
		{
			group.POST("/get_user_by_code", auth.MiniProgramLogin())
			group.POST("/bind_user_by_openid", auth.BindUser())
			group.POST("/login", auth.SelectUser())

			//use := group.Use(middleware.WechatMiniProgram(srv.C.Auth.OpenApi, srv.DbConnects))
			//{
			//	use.GET("/order/get", api.OrderList())
			//}
		}
	}
}
