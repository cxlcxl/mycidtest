package test

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/panjf2000/ants/v2"
	"gorm.io/gorm"
	"net/url"
	"path"
	"strconv"
	"testing"
	"time"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/mylog"
	"xiaoniuds.com/cid/vars"
)

var eventTypes = map[int8]string{1: "click", 2: "expose"}

func TestJcReport(t *testing.T) {
	sysLogPath := path.Join("../log", "syslog", time.Now().Format("20060102"))
	vars.SysLog = mylog.NewLog(sysLogPath)
	vars.Config, _ = config.LoadConfig("../config/config.yaml")

	db := data.NewDB()

	links, _ := common.NewJcLinkModel("", db).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("is_delete = 0 and main_user_id != 12000020828")
	}, []string{})
	linkMap := make(map[int64]int64)
	for _, link := range links {
		linkMap[link.JcConfigId] = link.ID
	}

	trackConfigs, _ := common.NewJcTrackConfigModel("", db).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Order("id desc")
	}, []string{})
	trackConfigMap := make(map[string]*common.JcTrackConfig)
	for _, trackConfig := range trackConfigs {
		trackConfigMap[fmt.Sprintf("%s_%d_%d", trackConfig.Pid, trackConfig.MediaType, 0)] = trackConfig
		if id, ok := linkMap[trackConfig.ID]; ok {
			trackConfigMap[fmt.Sprintf("%s_%d_%d", trackConfig.Pid, trackConfig.MediaType, id)] = trackConfig
		}
	}

	list, err := common.NewJcReportLogModel("", db).QueryByBuilder(func(db *gorm.DB) *gorm.DB {
		return db.Where("report_time >= UNIX_TIMESTAMP('2025-01-05 10:00:00')").Limit(1000).Offset(0)
	}, []string{
		"id", "ad_site_id", "event_type", "pid", "media_type", "udid", "oaid", "ip", "ua", "os",
		"click_time", "callback", "campaign_id", "ad_id", "creative_id", "advertiser_id", "log_key",
	})
	if err != nil {
		t.Fatal("失败", err)
	}

	pool, e := ants.NewPool(200)
	if e != nil {
		t.Fatal("失败", e)
	}
	defer pool.Release()

	resultChan := make(chan *ReportResult, 300)

	if len(list) > 0 {
		pool.Submit(func() {
			for result := range resultChan {
				t.Log("上报结果", result, "执行中的", pool.Running())
			}
		})
		for _, l := range list {
			configKey := fmt.Sprintf("%s_%d_%d", l.Pid, l.MediaType, l.AdSiteId)
			if trackConfig, ok := trackConfigMap[configKey]; ok {
				_ = pool.Submit(report(l, resultChan, trackConfig))
			}
		}

		close(resultChan)
	} else {
		t.Log("暂无数据")
	}

	pool.Waiting()
}

type ReportResult struct {
	LogId   int64
	Success bool
}

type ReportResponse struct {
	Success bool `json:"success"`
}

// {"resultCode":"SUCCESS","resultDesc":"成功","retriable":false,"success":true}
func report(reportLog *common.JcReportLog, resultChan chan *ReportResult, config *common.JcTrackConfig) func() {
	return func() {
		i := 1

		imei, idfa := "", ""
		if reportLog.Os == 0 {
			imei = reportLog.Udid
		} else {
			idfa = reportLog.Udid
		}
		requestParams := map[string]string{
			"action":      "click",
			"requestFrom": config.RequestFrom,
			"pid":         reportLog.Pid,
			"partnerId":   config.PartnerId,
			"benefit":     config.Benefit,
			"imei":        imei,
			"idfamd5":     idfa,
			"oaidmd5":     reportLog.Oaid,
			"ip":          reportLog.Ip,
			"ua":          reportLog.Ua,
			"os":          strconv.Itoa(int(reportLog.Os)),
			"timestamp":   strconv.Itoa(int(reportLog.ClickTime * 1000)),
			"callback":    getCallbackParam(reportLog),
			"accountid":   config.CallbackAccount,
			"campaignid":  strconv.Itoa(int(reportLog.CampaignId)),
			"adid":        strconv.Itoa(int(reportLog.AdId)),
			"cid":         strconv.Itoa(int(reportLog.CreativeId)),
		}

		for {
			resp, _ := resty.New().R().SetQueryParams(requestParams).Get("https://ugapi.alipay.com/monitor")

			var response ReportResponse
			_ = json.Unmarshal(resp.Body(), &response)

			res := &ReportResult{
				LogId:   reportLog.ID,
				Success: resp.StatusCode() == 200 && response.Success,
			}
			if res.Success || i >= 3 {
				resultChan <- res
				break
			}

			i++
		}
	}
}

func getCallbackParam(l *common.JcReportLog) string {
	callback := url.QueryEscape(fmt.Sprintf(
		"https://cid.xiaoniuds.com/Cid/Jc/callback?uk=%s&media_type=%d&pid=%s&click_time=%d&event_type=%s",
		l.LogKey, l.MediaType, l.Pid, l.ClickTime, eventTypes[l.EventType],
	))

	return callback
}
