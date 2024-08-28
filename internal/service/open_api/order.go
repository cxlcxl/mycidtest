package open_api

import (
	"xiaoniuds.com/cid/app/open_api/statement"
	"xiaoniuds.com/cid/config"
	"xiaoniuds.com/cid/internal/data"
)

type Order struct {
	C         *config.Config
	DbConnect *data.Data
}

func (o *Order) OrderList(params statement.OrderList) {

}
