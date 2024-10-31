package open_api

import (
	"fmt"
	"gorm.io/gorm"
	"strings"
	"time"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/data/base"
	"xiaoniuds.com/cid/internal/data/common"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/pkg/util"
	"xiaoniuds.com/cid/vars"
)

type PDD struct {
	PayTimeField    string
	VerifyTimeField string
}

func (p *PDD) GetOrderList(params statement.OrderList, connects *data.Data) (res interface{}, total int64, err *errs.MyErr) {
	sort := p.getSortField(params.TimeType)
	// 限制只能查看企业 main_user_id
	advWheres := []string{
		fmt.Sprintf("ad_act.main_user_id = %d", params.OpenApiData.MainUserId),
		"ad_act.no_callback_reason != '跨店订单不回传'",
		"ad_act.is_hide = 0",
	}
	if params.TimeType == TimeTypeCreateTime {
		advWheres = append(
			advWheres,
			fmt.Sprintf("ad_act.create_time >= '%s'", params.StartTime.Format(time.DateTime)),
			fmt.Sprintf("ad_act.create_time <= '%s'", params.EndTime.Format(time.DateTime)),
		)
	}
	if params.TimeType == TimeTypeUpdateTime {
		advWheres = append(
			advWheres,
			fmt.Sprintf("ad_act.update_time >= '%s'", params.StartTime.Format(time.DateTime)),
			fmt.Sprintf("ad_act.update_time <= '%s'", params.EndTime.Format(time.DateTime)),
		)
	}
	if params.TimeType == TimeTypePayTime {
		advWheres = append(
			advWheres,
			fmt.Sprintf("ad_act.%s >= %d", p.PayTimeField, params.StartTime.Unix()),
			fmt.Sprintf("ad_act.%s <= %d", p.PayTimeField, params.EndTime.Unix()),
		)
	}
	if params.TimeType == TimeTypeVerifyTime {
		advWheres = append(
			advWheres,
			fmt.Sprintf("ad_act.%s >= %d", p.VerifyTimeField, params.StartTime.Unix()),
			fmt.Sprintf("ad_act.%s <= %d", p.VerifyTimeField, params.EndTime.Unix()),
		)
	}
	baseFields := []string{
		"ad_act.id", "ad_act.order_sn", "ad_act.p_id", "ad_act.p_name", "ad_act.is_hide", "ad_act.goods_id",
		"ad_act.goods_name", "ad_act.goods_quantity", "ad_act.promotion_rate", "ad_act.goods_thumbnail_url", "ad_act.order_amount",
		"ad_act.mall_id", "ad_act.mall_name", "ad_act.order_status", "ad_act.order_status_desc", "ad_act.fail_reason",
		"ad_act.type", "ad_act.click_id", "ad_act.ad_id", "ad_act.ad_site_id", "ad_act.media_type", "ad_act.account_user_id",
		"ad_act.is_callback", "ad_act.no_callback_reason", "ad_act.cl_trace_type", "ad_act.is_direct", "ad_act.callback_type",
		"ad_act.owner_user_id", "ad_act.promotion_type", "ad_act.is_diff_goods", "ad_act.custom_parameters",
		"ad_act.csite_type", "ad_act.callback_event", "ad_act.main_user_id", "ad_act.create_time", "ad_act.update_time",
	}
	groupByFields := append(
		baseFields,
		"ad_act.promotion_amount",
		"ad_act.goods_price",
		"ad_act.advertiser_id,med.advertiser_nick",
		"ad_act.id",
	)
	// 字段异常问题
	selectFields := append(
		baseFields,
		"COALESCE(ad_act.promotion_amount, 0) as promotion_amount",
		"COALESCE(ad_act.goods_price, 0) as goods_price",
		"CASE WHEN COALESCE(max(ad_act.order_pay_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_pay_time)))  END as order_pay_time",
		"CASE WHEN COALESCE(max(ad_act.order_create_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_create_time)))  END as order_create_time",
		"CASE WHEN COALESCE(max(ad_act.order_verify_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_verify_time)))  END as order_verify_time",
		"CASE WHEN COALESCE(max(ad_act.order_receive_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.order_receive_time)))  END as order_receive_time",
		"CASE WHEN COALESCE(max(ad_act.callback_time),0) = 0 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.callback_time)))  END as callback_time",
		//order_modify_at字段下线，使用退款回传时间充当退款时间
		"CASE WHEN COALESCE(max(ad_act.order_status),0) != 4 THEN '--' ELSE COALESCE(max(FROM_UNIXTIME(ad_act.refund_callback_time)))  END as order_refund_time",
		"ad_act.advertiser_id,med.advertiser_nick",
		"MAX(t_admin_user.user_name) as user_name",
		"MAX(t_admin_user.user_full_name) as user_full_name",
	)
	leftJoins := []string{
		"LEFT JOIN chuangliang_doris_common.media_account AS med ON ad_act.advertiser_id = med.advertiser_id and ad_act.main_user_id = med.main_user_id and med.is_delete=0 and med.advertiser_id>0",
		"LEFT JOIN chuangliang_doris_common.admin_user t_admin_user ON ad_act.main_user_id = t_admin_user.main_user_id and t_admin_user.parent_id = 0",
	}

	selectFieldSQL := strings.Join(selectFields, ",")
	groupByFieldSQL := strings.Join(groupByFields, ",")
	advWhereSQL := strings.Join(advWheres, " AND ")
	leftJoinSQL := strings.Join(leftJoins, " ")

	offset := (params.Page - 1) * params.PageSize
	// id 辅助排序，防止大量 $sortField 重复导致乱序
	listSQL := fmt.Sprintf(
		"SELECT %s,COUNT() OVER() AS total_count FROM chuangliang_doris_cid.ad_order_pdd AS ad_act %s WHERE %s GROUP BY %s ORDER BY id ASC,%s DESC LIMIT %d OFFSET %d",
		selectFieldSQL, leftJoinSQL, advWhereSQL, groupByFieldSQL, sort, params.PageSize, offset,
	)
	var orders []*PddOrderItem
	e := data.NewDorisModel("", connects).QuerySQL(listSQL, &orders)
	if e != nil {
		err = errs.Err(errs.SysError, e)
	}
	if len(orders) > 0 {
		total = orders[0].TotalCount

		var (
			emptyMainUserIdPid []string
			emptyNamePid       []string
		)
		for _, order := range orders {
			if order.MainUserId == 0 {
				emptyMainUserIdPid = append(emptyMainUserIdPid, order.PID)
			}
			if order.PName == "" {
				emptyNamePid = append(emptyNamePid, order.PID)
			}
		}

		type adminUser struct {
			mainUserId   int64
			userFullName string
		}
		pidMap := make(map[string]*adminUser)
		namePidMaps := make(map[string]string)
		if len(emptyMainUserIdPid) > 0 {
			emptyMainUserIdPid = util.ArrayUnique(emptyMainUserIdPid)
			var pidList []*common.PddPidList
			pidList, err = common.NewPddPidListModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("pid in ?", emptyMainUserIdPid)
			}, []string{"pid", "main_user_id"})
			if err != nil {
				return
			}
			MainUserIds := make([]int64, len(pidList))
			for i, list := range pidList {
				MainUserIds[i] = list.MainUserId
			}
			MainUserIds = util.ArrayUnique(MainUserIds)

			var userList []*base.User
			userList, err = base.NewUserModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("user_id in ?", MainUserIds)
			}, []string{"user_full_name", "user_id"})
			if err != nil {
				return
			}
			userMap := map[int64]string{}
			for _, user := range userList {
				userMap[user.UserId] = user.UserFullName
			}

			for _, list := range pidList {
				pidMap[list.Pid] = &adminUser{
					mainUserId:   list.MainUserId,
					userFullName: userMap[list.MainUserId],
				}
			}
		}
		if len(emptyNamePid) > 0 {
			emptyNamePid = util.ArrayUnique(emptyNamePid)
			var pidList []*common.PddPidList
			pidList, err = common.NewPddPidListModel("", connects).GetListByBuilder(func(db *gorm.DB) *gorm.DB {
				return db.Where("pid in ?", emptyNamePid)
			}, []string{"pid", "pid_name"})
			if err != nil {
				return
			}
			for _, list := range pidList {
				namePidMaps[list.Pid] = list.PidName
			}
		}

		// 商品备注
		for i, order := range orders {
			if typeDesc, ok := vars.PddOrderType[order.Type]; ok {
				orders[i].TypeDesc = typeDesc
			} else {
				orders[i].TypeDesc = "其他"
			}
			if order.OrderRefundTime == "1970-01-01 08:00:00" {
				order.OrderRefundTime = "--"
			}
			if directDesc, ok := vars.OrderDirect[order.IsDirect]; ok {
				orders[i].IsDirectDesc = directDesc
			} else {
				orders[i].IsDirectDesc = "--"
			}
			if traceTypeMsg, ok := vars.TraceType[order.ClTraceType]; ok {
				orders[i].TraceTypeMsg = traceTypeMsg
			} else {
				orders[i].TraceTypeMsg = "--"
			}
			orders[i].CSiteTypeDesc = vars.AdSiteType[order.CSiteType]            // 广告版位
			orders[i].CallbackEventDesc = vars.CallbackEvent[order.CallbackEvent] // 回传事件

			if order.MainUserId == 0 {
				if pid, ok := pidMap[order.PID]; ok {
					orders[i].MainUserId = pid.mainUserId
					orders[i].UserFullName = pid.userFullName
				}
			}
			if order.PName == "" {
				orders[i].PName = namePidMaps[order.PID]
			}
		}
	}
	res = orders

	return
}

func (p *PDD) getSortField(timeType int) (field string) {
	switch timeType {
	case TimeTypeCreateTime:
		field = "create_time"
		break
	case TimeTypeUpdateTime:
		field = "update_time"
	case TimeTypePayTime:
		field = p.PayTimeField
		break
	case TimeTypeVerifyTime:
		field = p.VerifyTimeField
		break
	default:
		field = "order_create_time"
	}
	return
}
