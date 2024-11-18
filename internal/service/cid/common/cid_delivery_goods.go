package common

import (
	"fmt"
	"gorm.io/gorm"
	"slices"
	"strconv"
	"strings"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type CidDeliveryGoodsService struct {
	DbConnect *data.Data
}

// SearchGoodsByAdvertisers 商品列表通过关联投放账户搜索
func (s *CidDeliveryGoodsService) SearchGoodsByAdvertisers(mainUserId int64, mediaType vars.Media, platform vars.Platform, actKeyword string) (goodsIds []string) {
	accounts := strings.Split(actKeyword, ",")
	values := make([]interface{}, 0)
	ids := make([]int64, 0)
	wheres := make([]string, 0)
	leftJoins := make([]string, 0)
	for _, account := range accounts {
		id, _ := strconv.ParseInt(account, 10, 64)
		if id == 0 {
			values = append(values, "%"+account+"%")
			wheres = append(wheres, fmt.Sprintf("t1.advertiser_nick like ?"))
		} else {
			ids = append(ids, id)
		}
	}

	searchWhere := ""
	if len(ids) > 0 {
		values = append(values, ids)
		searchWhere += " and t0.advertiser_id in ?"
	}
	// 有名称的搜索需要连表
	if len(wheres) > 0 {
		leftJoins = append(leftJoins, "left join chuangliang_doris_common.media_account as t1 on t0.main_user_id = t1.main_user_id and t0.advertiser_id = t1.advertiser_id and t1.is_delete = 0")
		searchWhere += " and (" + strings.Join(wheres, " or ") + ")"
	}
	leftJoinSQL := strings.Join(leftJoins, " ")
	goodIdField := "t0.goods_id"
	if platform == vars.PlatformJd {
		goodIdField = "if(ifnull(t0.sku_id, 0) > 0, t0.sku_id, t0.goods_id) as goods_id"
	}
	sql := fmt.Sprintf("select %s from chuangliang_doris_cid.cid_delivery_goods as t0 %s where t0.`is_delete` = 0 and t0.`media_type` = %d and t0.`platform` = %d and t0.`main_user_id` = %d %s", goodIdField, leftJoinSQL, mediaType, platform, mainUserId, searchWhere)
	_ = data.NewDorisModel("", s.DbConnect).QuerySQL(sql, &goodsIds, values...)
	if len(goodsIds) > 0 {
		goodsIds = util.ArrayUnique(goodsIds)
	} else {
		goodsIds = []string{"-1"}
	}

	return nil
}

type BelongsGoodsInfo struct {
	AdvertiserId   int64  `json:"advertiser_id"`
	AdvertiserNick string `json:"advertiser_nick"`
	GoodsId        string `json:"goods_id"`
	MainUserId     int64  `json:"main_user_id"`
	MediaType      string `json:"media_type"`
}

// BelongsTopAccountsByGoodsIds 列表关联订单量第一的账户
func (s *CidDeliveryGoodsService) BelongsTopAccountsByGoodsIds(mainUserId int64, platform vars.Platform, goodsIds []int64) (belongsMap map[string]*BelongsGoodsInfo) {
	belongsMap = make(map[string]*BelongsGoodsInfo)
	if len(goodsIds) == 0 {
		return
	}
	goodsIds = util.ArrayUnique(goodsIds)
	if !slices.Contains([]vars.Platform{vars.PlatformTb, vars.PlatformPdd, vars.PlatformJd}, platform) {
		return
	}
	deliveryGoods, _ := common.NewCidDeliveryGoodsModel("", s.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("main_user_id = ?", mainUserId).
			Where("is_delete = 0").
			Where("goods_id in ?", goodsIds).
			Order("goods_id asc,order_cnt_30d desc,delivery_date desc,id desc")
	}, []string{"advertiser_id", "goods_id", "main_user_id", "media_type"})
	if len(deliveryGoods) == 0 {
		return
	}
	var (
		deliveryAccountFirst = make([]*BelongsGoodsInfo, 0)
		advertiserIds        = make([]int64, 0)
		exists               = make(map[string]struct{})
	)
	for _, goods := range deliveryGoods {
		// 只需要第一个
		if _, ok := exists[goods.GoodsId]; ok {
			continue
		}
		exists[goods.GoodsId] = struct{}{}
		deliveryAccountFirst = append(deliveryAccountFirst, &BelongsGoodsInfo{
			AdvertiserId: goods.AdvertiserId,
			GoodsId:      goods.GoodsId,
			MainUserId:   mainUserId,
			MediaType:    vars.MediaTypeInt2Str[vars.Media(goods.MediaType)],
		})
		advertiserIds = append(advertiserIds, goods.AdvertiserId)
	}

	accounts, _ := base.NewMediaAccountModel("", s.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("main_user_id = ?", mainUserId).
			Where("is_delete = 0").
			Where("advertiser_id in ?", util.ArrayUnique(advertiserIds))
	}, []string{"advertiser_id", "advertiser_nick"})
	accountMap := make(map[int64]string)
	for _, account := range accounts {
		accountMap[account.AdvertiserId] = account.AdvertiserNick
	}
	for i, goods := range deliveryAccountFirst {
		deliveryAccountFirst[i].AdvertiserNick = accountMap[goods.AdvertiserId]

		belongsMap[goods.GoodsId] = deliveryAccountFirst[i]
	}
	return
}
