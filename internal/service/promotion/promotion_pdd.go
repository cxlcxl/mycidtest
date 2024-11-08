package promotion

import (
	"fmt"
	"github.com/jinzhu/copier"
	"gorm.io/gorm"
	"slices"
	"strconv"
	"strings"
	"time"
	"xiaoniuds.com/cid/app/cid/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/internal/data/common"
	common2 "xiaoniuds.com/cid/internal/service/common"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type PddGoods struct {
	C         *config.Config
	DbConnect *data.Data
}

type PddGoodsItem struct {
	*common.PddGoods
	DeliveryAdvertiserId   int64    `json:"delivery_advertiser_id"`
	DeliveryAdvertiserNick string   `json:"delivery_advertiser_nick"`
	DeliveryMediaType      string   `json:"delivery_media_type"`
	IsBelong               string   `json:"is_belong"`
	MallCertificateUrls    []string `json:"mall_certificate_urls"`
	ThumbPicUrls           []string `json:"thumb_pic_urls"`
	Note                   string   `json:"note"`
	OwnerUserName          string   `json:"owner_user_name"`
}

type PddGoodsSale struct {
	GoodsId      int64 `json:"goods_id"`
	SaleNum      int64 `json:"sale_num"`
	TodaySaleNum int64 `json:"today_sale_num"`
}

func (g *PddGoods) List(params statement.PddGoodsList) (goods []*PddGoodsItem, total int64, err *errs.MyErr) {
	offset := (params.Page - 1) * params.PageSize
	if params.SortField == "" {
		params.SortField = "create_time"
	}
	if params.SortDirection == "" {
		params.SortDirection = "desc"
	}
	var (
		DeliveryMedia      vars.Media = 0
		DeliveryAdvertiser            = ""
		goodsIds                      = make([]string, 0)
		pddGoodsPrimaryIds            = make([]int64, 0)
	)
	before30Day := time.Now().AddDate(0, 0, -30)
	if params.TTAdvertiser != "" {
		DeliveryMedia = vars.MediaTypeIntTT
		DeliveryAdvertiser = params.TTAdvertiser
	}
	if params.KSAdvertiser != "" {
		DeliveryMedia = vars.MediaTypeIntKS
		DeliveryAdvertiser = params.KSAdvertiser
	}
	if params.GDTAdvertiser != "" {
		DeliveryMedia = vars.MediaTypeIntGDT
		DeliveryAdvertiser = params.GDTAdvertiser
	}
	if DeliveryMedia > 0 {
		goodsIds = (&common2.CidDeliveryGoodsService{
			C:         g.C,
			DbConnect: g.DbConnect,
		}).SearchGoodsByAdvertisers(
			params.LoginData.MainUserId,
			DeliveryMedia,
			vars.PlatformPdd,
			DeliveryAdvertiser,
		)
	}

	goodsNoteMap := make(map[int64]string)
	// 开启只查询我的备注 或者 传入备注模糊查询
	if params.Note != "" || params.OwnerNoteStatus == 1 {
		notes, _ := common.NewPddGoodsNoteRelModel("", g.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
			if params.Note != "" {
				db = db.Where("note like ?", "%"+strings.TrimSpace(params.Note)+"%")
			}
			return db.Where("main_user_id = ?", params.LoginData.MainUserId).
				Where("owner_user_id = ?", params.LoginData.UserId)
		}, []string{"record_id", "note"})
		if len(notes) == 0 {
			pddGoodsPrimaryIds = []int64{-1}
		} else {
			for _, note := range notes {
				pddGoodsPrimaryIds = append(pddGoodsPrimaryIds, note.RecordId)
				goodsNoteMap[note.RecordId] = note.Note
			}
		}
	}
	builder := func(db *gorm.DB) *gorm.DB {
		if slices.Contains([]string{"sale_num", "today_sale_num"}, params.SortField) {
			today := time.Now().Format(time.DateOnly)
			db.Joins(fmt.Sprintf(
				"left join (select goods_id,sum(sale_num) as sale_num,sum(if(stat_date = '%s', sale_num, 0)) as today_num "+
					"from goods_sale_sum where main_user_id = %d and stat_date >= '%s' and platform = %d "+
					"group by goods_id) as sale_tbl on sale_tbl.goods_id = pdd_goods.goods_id",
				today, params.LoginData.MainUserId, before30Day.Format(time.DateTime), vars.PlatformPdd,
			)).Select("pdd_goods.*", "IFNULL(sale_tbl.sale_num, 0) as sale_num", "IFNULL(sale_tbl.today_num, 0) as today_sale_num")
		}

		if params.DeliveryStatus >= 0 {
			db = db.Where("pdd_goods.delivery_status = ?", params.DeliveryStatus)
		}
		if len(goodsIds) > 0 {
			db = db.Where("pdd_goods.goods_id in ?", goodsIds)
		}
		if params.ReportStatus >= 0 && params.ReportStatus <= 5 {
			db = db.Where("pdd_goods.report_status = ?", params.ReportStatus)
		}
		if params.StartDate != "" {
			db = db.Where("pdd_goods.create_time >= ?", params.StartDate+" 00:00:00")
		}
		if params.EndDate != "" {
			db = db.Where("pdd_goods.create_time <= ?", params.EndDate+" 23:59:59")
		}
		if len(params.MallIds) > 0 {
			db = db.Where("pdd_goods.mall_id in ?", params.MallIds)
		}
		if len(pddGoodsPrimaryIds) > 0 {
			db = db.Where("pdd_goods.id in ?", pddGoodsPrimaryIds)
		}
		return db.Where("pdd_goods.main_user_id = ?", params.LoginData.MainUserId).
			Order(fmt.Sprintf("%s %s", params.SortField, params.SortDirection))
	}
	//        // 数据权限
	//        $ownerUserIds = \App\Services\Cid\CidDataPermissionService::getShowRange(static::MODULE_NAME);
	//        $conditions['owner_user_id'] = $ownerUserIds;
	//        if(!empty($params['owner_user_id'])){ //筛选所属人
	//            $conditions['owner_user_id'] = array_intersect($params['owner_user_id'], $conditions['owner_user_id']);
	//        }
	//        if (!empty($params['keyword'])){
	//            $conditions['keyword'] = $params['keyword'];
	//        }

	list, total, err := common.NewPddGoodsModel("", g.DbConnect).QueryListByBuilder(builder, []string{}, offset, params.PageSize)
	if len(list) > 0 {
		e := copier.Copy(&goods, &list)
		if e != nil {
			err = errs.Err(errs.SysError, e)
			return
		}
		var (
			adminUserMap        = make(map[int64]string)
			pddGoodsOrderMap    = make(map[int64]*PddGoodsSale)
			deliveryAccountsMap = make(map[string]*common2.BelongsGoodsInfo)
			userIds             = make([]int64, 0)
			ids                 = make([]int64, 0)
			pddGoodsIds         = make([]int64, 0)
		)
		// 创建人
		for _, pddGoods := range goods {
			ids = append(ids, pddGoods.ID)
			userIds = append(userIds, pddGoods.OwnerUserId)
			pddGoodsIds = append(pddGoodsIds, pddGoods.GoodsId)
		}
		userIds = util.ArrayUnique(userIds)
		users, _ := base.NewUserModel("", g.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
			return db.Where("user_id in ?", userIds)
		}, []string{"user_id", "user_name"})
		for _, user := range users {
			adminUserMap[user.UserId] = user.UserName
		}
		// 商品ID，获取近30天销量
		if slices.Contains([]string{"sale_num", "today_sale_num"}, params.SortField) {
			t := time.Now()
			today := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
			adOrderSQL := fmt.Sprintf("SELECT goods_id,COUNT(1) AS sale_num,SUM(IF(order_pay_time >= %d,1,0)) AS today_sale_num FROM chuangliang_doris_cid.ad_order_pdd WHERE main_user_id = ? AND order_pay_time >= ? AND goods_id IN ? GROUP BY goods_id", today.Unix())
			var goodsSales []*PddGoodsSale
			_ = data.NewDorisModel("", g.DbConnect).QuerySQL(adOrderSQL, &goodsSales, params.LoginData.MainUserId, before30Day.Unix(), pddGoodsIds)
			for _, goodsSale := range goodsSales {
				pddGoodsOrderMap[goodsSale.GoodsId] = goodsSale
			}
		}

		if len(goodsNoteMap) == 0 {
			notes, _ := common.NewPddGoodsNoteRelModel("", g.DbConnect).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("main_user_id = ?", params.LoginData.MainUserId).
					Where("owner_user_id = ?", params.LoginData.UserId).
					Where("record_id in ?", ids)
			}, []string{"record_id", "note"})
			for _, note := range notes {
				goodsNoteMap[note.RecordId] = note.Note
			}
		}

		deliveryAccountsMap = (&common2.CidDeliveryGoodsService{C: g.C, DbConnect: g.DbConnect}).BelongsTopAccountsByGoodsIds(params.LoginData.MainUserId, vars.PlatformPdd, pddGoodsIds)
		for i, pddGoods := range goods {
			// 账号所属判断
			goods[i].IsBelong = "main"
			// 主账户是自己，或者所属人是自己
			if slices.Contains([]int64{pddGoods.OwnerUserId, params.LoginData.MainUserId}, params.LoginData.UserId) {
				goods[i].IsBelong = "self"
			}
			if pddGoods.OwnerUserId != params.LoginData.UserId &&
				slices.Contains([]string{"企业管理员"}, params.LoginData.GroupName) &&
				params.OwnerUserId != params.LoginData.MainUserId {
				goods[i].IsBelong = "child"
			}
			if slices.Contains(g.C.DuoIds, pddGoods.DuoId) {
				goods[i].IsBelong = "sys"
			}
			stringGoodsId := strconv.FormatInt(pddGoods.GoodsId, 10)
			if actMap, ok := deliveryAccountsMap[stringGoodsId]; ok {
				goods[i].DeliveryAdvertiserId = actMap.AdvertiserId
				goods[i].DeliveryAdvertiserNick = actMap.AdvertiserNick
				goods[i].DeliveryMediaType = actMap.MediaType
			} else {
				goods[i].DeliveryAdvertiserId = 0
				goods[i].DeliveryAdvertiserNick = "--"
				goods[i].DeliveryMediaType = "--"
			}

			if pddGoods.MallCertificateUrl != "" {
				goods[i].MallCertificateUrls = strings.Split(pddGoods.MallCertificateUrl, ",")
			} else {
				goods[i].MallCertificateUrls = []string{}
			}
			if pddGoods.ThumbPicUrl != "" {
				goods[i].ThumbPicUrls = strings.Split(pddGoods.ThumbPicUrl, ",")
			} else {
				goods[i].ThumbPicUrls = []string{}
			}
			if note, ok := goodsNoteMap[pddGoods.ID]; ok {
				goods[i].Note = note
			}
			if ownerUserName, ok := adminUserMap[pddGoods.OwnerUserId]; ok {
				goods[i].OwnerUserName = ownerUserName
			} else {
				goods[i].OwnerUserName = "--"
			}
			if pddGoods.DemoUrl == "" {
				goods[i].DemoUrl = goods[i].ImageVideoUrl
			}
			if !slices.Contains([]string{"sale_num", "today_sale_num"}, params.SortField) {
				goods[i].SaleNum = 0
				goods[i].TodaySaleNum = 0
				if sale, ok := pddGoodsOrderMap[pddGoods.GoodsId]; ok {
					goods[i].SaleNum = sale.SaleNum
					goods[i].TodaySaleNum = sale.TodaySaleNum
				}
			}
		}
	}

	return
}
