package task

import (
	"xiaoniuds.com/cid/pkg/auth_token"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"

	"gorm.io/gorm"
	"xiaoniuds.com/cid/internal/data"
)

type TaskLogDownloadCenter struct {
	TaskId         int64            `json:"task_id"`          // 任务ID
	ParentTaskId   int64            `json:"parent_task_id"`   // 父任务ID
	RootTaskId     int64            `json:"root_task_id"`     // 根节点任务ID
	BatchId        string           `json:"batch_id"`         // 生成任务批次唯一标识
	WorkerId       int64            `json:"worker_id"`        // worker_id
	QueueName      string           `json:"queue_name"`       //
	MediaAccountId int64            `json:"media_account_id"` // 媒体账号ID
	Status         string           `json:"status"`           // 任务状态
	TaskDate       *data.DbDate     `json:"task_date"`        // 任务日期
	TaskRoute      string           `json:"task_route"`       // 任务名称
	TaskParam      string           `json:"task_param"`       // 任务参数
	TaskParamMd5   string           `json:"task_param_md5"`   // 任务参数md5
	Priority       int64            `json:"priority"`         // 优先级，数字越大越优先
	CostTime       int64            `json:"cost_time"`        // 花费时间(毫秒)
	ResultData     string           `json:"result_data"`      // 结果数据
	CreateTime     *data.DbDateTime `json:"create_time"`      // 创建时间
	UpdateTime     *data.DbDateTime `json:"update_time"`      // 最后更新时间
	RetryNum       int8             `json:"retry_num"`        //
	TaskName       string           `json:"task_name"`        // 任务名称
	TaskStatus     int8             `json:"task_status"`      // 任务状态： 0: 处理中 1：处理完成 2: 处理失败
	RequestParams  string           `json:"request_params"`   // 请求参数
	DownloadUrl    string           `json:"download_url"`     // 下载URL
	CreateUserId   int64            `json:"create_user_id"`   //
	Type           int8             `json:"type"`             // 报表类型 1: 广告报表 2：素材报表
	UpdateUserId   int64            `json:"update_user_id"`   //
	IsDelete       int8             `json:"is_delete"`        // 是否已删除
	Msg            string           `json:"msg"`              // 导出错误信息
	MediaType      string           `json:"media_type"`       // 媒体类型
	ProductVersion int8             `json:"product_version"`  // 产品版本
	ID             int64            `json:"id"`               //
	MainUserId     int64            `json:"main_user_id"`     // 租户id
	EnvType        int16            `json:"env_type"`         // [0,1]是预发布脚本机
	SystemParam    string           `json:"system_param"`     // 系统参数扩展
}

type DownloadCenterModel struct {
	dbName string
	db     *gorm.DB
}

func NewDownloadCenterModel(connect string, connects *data.Data) *DownloadCenterModel {
	if connect == "" {
		connect = "ad_task"
	}
	return &DownloadCenterModel{
		dbName: "task_log_download_center",
		db:     connects.DbConnects[connect],
	}
}

type DownloadCenterListItem struct {
	ID          int64            `json:"id"`
	TaskName    string           `json:"task_name"`
	TaskStatus  int8             `json:"task_status"`
	DownloadUrl string           `json:"download_url"`
	CreateTime  *data.DbDateTime `json:"create_time"`
}

func (m *DownloadCenterModel) GetDownloadCenterList(
	loginUser *auth_token.LoginData,
	taskName string, taskType, page, pageSize int,
) (
	downloadLogs []*DownloadCenterListItem, total int64, err *errs.MyErr,
) {
	query := m.db.Debug().Table(m.dbName).
		Select("id", "task_name", "task_status", "download_url", "create_time").
		Where("create_user_id = ?", loginUser.UserId).
		Where("product_version = ?", loginUser.ProductVersion).
		Where("is_delete = 0")
	if taskName != "" {
		query = query.Where("task_name LIKE ?", "%"+taskName+"%")
	}
	if taskType > 0 {
		query = query.Where("type = ?", taskType)
	}
	if query.Count(&total); total == 0 {
		return
	}

	e := query.Offset(util.Offset(page, pageSize)).Limit(pageSize).
		Find(&downloadLogs).Error
	if e != nil {
		return nil, 0, errs.Err(errs.SysError, e)
	}

	return
}
