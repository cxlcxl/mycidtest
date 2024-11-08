package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type CidDeliveryGoods struct {
	ID           int64       `json:"id"`
	MainUserId   int64       `json:"main_user_id"`  //
	AdvertiserId int64       `json:"advertiser_id"` // 媒体平台的账号ID
	MediaType    int         `json:"media_type"`    // 媒体类型
	Platform     int         `json:"platform"`      // 电商平台
	GoodsId      string      `json:"goods_id"`      //
	GoodsName    string      `json:"goods_name"`    //
	PreviewUrl   string      `json:"preview_url"`   //
	DeliveryDate data.DbDate `json:"delivery_date"` // 投放日期
	OrderCnt30d  int         `json:"order_cnt_30d"` // 30天出单量
	DataFrom     int8        `json:"data_from"`     // 数据来源
	SkuId        int64       `json:"sku_id"`
	CreateUserId int64       `json:"create_user_id"`
	IsDelete     uint8       `json:"is_delete"`
	*data.Timestamp
}

type CidDeliveryGoodsModel struct {
	dbName string
	db     *gorm.DB
}

func NewCidDeliveryGoodsModel(connect string, connects *data.Data) *CidDeliveryGoodsModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &CidDeliveryGoodsModel{
		dbName: "cid_delivery_goods",
		db:     connects.DbConnects[connect],
	}
}

func (m *CidDeliveryGoodsModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*CidDeliveryGoods, err *errs.MyErr) {
	query := m.db.Table(m.dbName).Where("is_delete = ?", 0)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	e := query.Find(&list).Error
	if e != nil {
		return nil, errs.Err(errs.SysError, e)
	}
	return
}
