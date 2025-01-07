package common

import (
	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type JcReportLog struct {
	ID           int64  `json:"id"`
	AdSiteId     int64  `json:"ad_site_id"`
	Pid          string `json:"pid"`         // 推广位字符串ID
	EventType    int8   `json:"event_type"`  // 1点击，2曝光
	ReportTime   int64  `json:"report_time"` // 上报时间
	ReportNum    int64  `json:"report_num"`
	MediaType    int8   `json:"media_type"` // 媒体类型
	Udid         string `json:"udid"`       // 设备ID，ios是idfa，安卓是imei
	Oaid         string `json:"oaid"`       // Android Q及更高版本的设备号
	Callback     string `json:"callback"`   // callback也就是我们说的click_id
	ClickTime    int64  `json:"click_time"` // 点击时间，落地页取服务器接收到时间，监控链接取传入的时间
	Os           int8   `json:"os"`         // 系统，0安卓，1ios，3其他
	Ip           string `json:"ip"`         // 客户的IP
	Ua           string `json:"ua"`
	CampaignId   int64  `json:"campaign_id"`   // 广告组ID
	AdId         int64  `json:"ad_id"`         // 广告ID
	CreativeId   int64  `json:"creative_id"`   // 创意ID
	AdvertiserId int64  `json:"advertiser_id"` // 媒体账户ID
	State        int8   `json:"state"`         // 上报状态：1表示成功
	LogKey       string `json:"log_key"`       // callback组合的key
}

// TableName 表名称
func (*JcReportLog) TableName() string {
	return "jc_report_log"
}

type JcReportLogModel struct {
	dbName string
	db     *gorm.DB
}

func NewJcReportLogModel(connect string, connects *data.Data) *JcReportLogModel {
	if connect == "" {
		connect = vars.DRCLCidAdCommonWrite
	}
	return &JcReportLogModel{
		dbName: "jc_report_log",
		db:     connects.DbConnects[connect],
	}
}

func (m *JcReportLogModel) QueryByBuilder(builder data.QueryBuilder, fields []string) (list []*JcReportLog, err *errs.MyErr) {
	query := m.db.Debug().Table(m.dbName)
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
