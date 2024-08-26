package validator

import (
	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/pkg/errs"
)

type ValidOpt func(interface{})

func BindJsonData(ctx *gin.Context, data interface{}, opts ...ValidOpt) *errs.MyErr {
	e := ctx.BindJSON(data)
	if e != nil {
		return errs.Err(errs.ParamError, e)
	}
	for _, opt := range opts {
		opt(data)
	}

	return nil
}
