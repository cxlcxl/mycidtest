package server

import (
	"xiaoniuds.com/cid/app/cid/handle/promotion"
	"xiaoniuds.com/cid/internal/middleware"
)

func NewPromotionServer() Opt {
	return func(srv *Server) {
		promotionApi := &promotion.Promotion{C: srv.C, DbConnect: srv.DbConnects}
		mediaActApi := &promotion.MediaAccount{C: srv.C, DbConnect: srv.DbConnects}

		goodsLink := srv.engine.Group("/goods", middleware.LoginAuth(srv.C.Auth.Login, srv.DbConnects))
		{
			goodsLink.GET("/pdd", promotionApi.PddGoods())
			goodsLink.GET("/tb", promotionApi.TbGoods())
			goodsLink.GET("/jd", promotionApi.JdGoods())
		}

		mediaAct := srv.engine.Group("/advertiser", middleware.LoginAuth(srv.C.Auth.Login, srv.DbConnects))
		{
			mediaAct.GET("/tt", mediaActApi.TtActList())
			mediaAct.GET("/ks", mediaActApi.KsActList())
			mediaAct.GET("/gdt", mediaActApi.GdtActList())
		}
	}
}
