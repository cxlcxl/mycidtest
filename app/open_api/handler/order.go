package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
	"xiaoniuds.com/cid/internal/service/open_api"
	"xiaoniuds.com/cid/pkg/util/response"
	"xiaoniuds.com/cid/pkg/util/validator"
)

type OpenOrder struct {
	C         *config.Config
	DbConnect *data.Data
}

func (o *OpenOrder) OrderList() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var params statement.OrderList
		if err := validator.BindJsonData(ctx, &params, func(d interface{}) error {
			end := d.(*statement.OrderList).EndTime.AddDate(0, 0, -1)
			if d.(*statement.OrderList).StartTime.Before(end) {
				return errors.New("最大查询 1 天的数据")
			}
			return nil
		}); err != nil {
			response.Error(ctx, err)
			return
		}

		//sortField := "order_create_time"
		//shopPayTimeField := [1 => "order_pay_time", 2 => "payment_time", 3 => "tk_paid_time"]
		//verifyTimeField := [1 => "order_verify_time", 2 => "payment_time", 3 => "tk_paid_time"]
		//$conditions = [];
		//switch ($timeType) {
		//case 1: //创建时间
		//$conditions['create_time'] = "$startTime,$endTime";
		//$sortField = "create_time";
		//break;
		//case 2://更新时间
		//$conditions['update_time'] = "$startTime,$endTime";
		//$sortField = "update_time";
		//break;
		//case 3://支付时间
		//$conditions[$shopPayTimeField[$shopType]] = "$startTime,$endTime";
		//$sortField = $shopPayTimeField[$shopType];
		//break;
		//case 4://审核时间
		//$conditions[$verifyTimeField[$shopType]] = "$startTime,$endTime";
		//$sortField = $verifyTimeField[$shopType];
		//}
		//$conditions['isDirect'] = $this->filterInput($request, "is_direct", 0);
		//$isHidden = (int)$request->input("is_hidden", -1);
		//if ($isHidden >= 0) {
		//$conditions['isHidden'] = $isHidden;
		//}
		//$conditions['pid'] = $this->filterInput($request, "pid", 0);
		//$conditions['id'] = $this->filterInput($request, "id", 0);
		//$conditions['mainUserIds'] = $this->filterInput($request, "main_user_ids", "");
		//$conditions['type'] = $this->filterInput($request, "type", -1);

		orders, total, err := (&open_api.Order{C: o.C, DbConnect: o.DbConnect}).OrderList(params)
		if err != nil {
			response.Error(ctx, err)
			return
		}

		response.PageSuccess(ctx, params.Page, params.PageSize, total, orders)
	}
}
