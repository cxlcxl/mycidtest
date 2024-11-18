package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type PddGoods struct {
	ID                 int64   `json:"id"`
	MainUserId         int64   `json:"main_user_id"`
	OwnerUserId        int64   `json:"owner_user_id"`
	DuoId              int64   `json:"duo_id"`               // 申请报备拼多多会员id
	GoodsName          string  `json:"goods_name"`           // 商品名称
	GoodsSign          string  `json:"goods_sign"`           // 拼多多商品签名值
	GoodsId            int64   `json:"goods_id"`             // 商品id
	GoodsImageUrl      string  `json:"goods_image_url"`      // 多多进宝商品主图
	GoodsThumbnailUrl  string  `json:"goods_thumbnail_url"`  // 商品缩略图
	MallId             int64   `json:"mall_id"`              // 店铺id
	MallName           string  `json:"mall_name"`            // 店铺名称
	DemoUrl            string  `json:"demo_url"`             // 推广商品视频素材url
	MallCertificateUrl string  `json:"mall_certificate_url"` // 商家资质证明图片url列表，1到3张图，用英文逗号隔开
	PromotionCodeUrl   string  `json:"promotion_code_url"`   // 推广视频预览码url
	ThumbPicUrl        string  `json:"thumb_pic_url"`        // 商品图片素材url列表，0到3张图，用英文逗号隔开
	PromotionStartTime int64   `json:"promotion_start_time"` // 推广结束时间戳
	PromotionEndTime   int64   `json:"promotion_end_time"`   // 推广开始时间戳
	CommitTime         int64   `json:"commit_time"`          // 报备提交时间
	IsPromotionLimited int8    `json:"is_promotion_limited"` // 是否被限制推广
	ReportStatus       int8    `json:"report_status"`        // 商品报备状态,0未报备,1报备中,2报备成功,3报备失败
	ReportStatusDesc   string  `json:"report_status_desc"`   // 商品报备状态,0未报备,1报备中,2报备成功,3报备失败
	ReportFailReason   string  `json:"report_fail_reason"`   // 报备失败原因
	MinGroupPrice      float64 `json:"min_group_price"`      // 最小拼团价
	SinglePriceStatus  int8    `json:"single_price_status"`  // 按单件价格回传，0未开启，1开启
	FixSettingType     int8    `json:"fix_setting_type"`     // 设置类型，1固定比例，2固定金额，3单件金额
	FixPriceSetting    float64 `json:"fix_price_setting"`    // 固定金额回传，0为未开启，其他为真实数值
	FixRateSetting     int     `json:"fix_rate_setting"`     // 按固定比例回传，0为未开启，其他为百分比数值
	CatIds             string  `json:"cat_ids"`              // 商品类目id
	SalesTip           string  `json:"sales_tip"`            // 商品销量
	ImageVideoUrl      string  `json:"image_video_url"`      // 图片视频
	IsDelete           uint8   `json:"is_delete"`
	DeliveryStatus     int8    `json:"delivery_status"`
	SaleNum            int64   `json:"sale_num"`
	TodaySaleNum       int64   `json:"today_sale_num"`
	*data.Timestamp
}

type PddGoodsModel struct {
	dbName string
	db     *gorm.DB
}

func NewPddGoodsModel(connect string, connects *data.Data) *PddGoodsModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommon
	}
	return &PddGoodsModel{
		dbName: "pdd_goods",
		db:     connects.DbConnects[connect],
	}
}

func (m *PddGoodsModel) QueryListByBuilder(builder data.QueryBuilder, fields []string, offset, pageSize int) (list []*PddGoods, total int64, err *errs.MyErr) {
	query := m.db.Debug().Table(m.dbName).Where("pdd_goods.is_delete = ?", 0)
	if len(fields) > 0 {
		query = query.Select(fields)
	}
	query = builder(query)
	_ = query.Count(&total).Error
	if total == 0 {
		return
	}
	e := query.Offset(offset).Limit(pageSize).Find(&list).Error
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	return
}
