package validator

import (
	"errors"
	"reflect"

	"github.com/gin-gonic/gin"
	"xiaoniuds.com/cid/pkg/errs"
	"xiaoniuds.com/cid/vars"
)

type ValidOpt func(interface{}) error

func BindJsonData(ctx *gin.Context, data interface{}, opts ...ValidOpt) *errs.MyErr {
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return errs.Err(errs.ParamError)
	}
	e := ctx.ShouldBind(data)
	if e != nil {
		return errs.Err(errs.ParamError, e)
	}

	for _, opt := range opts {
		if e = opt(data); e != nil {
			return errs.Err(errs.ParamError, e)
		}
	}

	if err := bindLoginUser(ctx, data); err != nil {
		return errs.Err(errs.ParamError, err)
	}
	if err := bindOpenApiLoginUser(ctx, data); err != nil {
		return errs.Err(errs.ParamError, err)
	}

	return nil
}

func bindLoginUser(ctx *gin.Context, data interface{}) *errs.MyErr {
	if _, ok := reflect.TypeOf(data).Elem().FieldByName("LoginData"); !ok {
		return nil
	}
	loginInfo, exists := ctx.Get(vars.LoginKey)
	if !exists {
		return errs.Err(errs.ParamError, errors.New("login info not exists"))
	}

	reflect.ValueOf(data).Elem().FieldByName("LoginData").Set(reflect.ValueOf(loginInfo))
	return nil
}

func bindOpenApiLoginUser(ctx *gin.Context, data interface{}) *errs.MyErr {
	if _, ok := reflect.TypeOf(data).Elem().FieldByName("OpenApiLoginData"); !ok {
		return nil
	}
	loginInfo, exists := ctx.Get(vars.OpenApiLoginKey)
	if !exists {
		return errs.Err(errs.ParamError, errors.New("login info not exists"))
	}

	reflect.ValueOf(data).Elem().FieldByName("OpenApiLoginData").Set(reflect.ValueOf(loginInfo))
	return nil
}
