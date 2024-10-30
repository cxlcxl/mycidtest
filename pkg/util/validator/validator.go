package validator

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"time"
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
	if _, ok := reflect.TypeOf(data).Elem().FieldByName("OpenApiData"); !ok {
		return nil
	}
	loginInfo, exists := ctx.Get(vars.OpenApiLoginKey)
	if !exists {
		return errs.Err(errs.ParamError, errors.New("login info not exists"))
	}

	reflect.ValueOf(data).Elem().FieldByName("OpenApiData").Set(reflect.ValueOf(loginInfo))
	return nil
}

// RegisterValidators 应用初始化时绑定
func RegisterValidators() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		for m, vFunc := range validators {
			_ = v.RegisterValidation(m, vFunc)
		}
	}
}

var validators = map[string]validator.Func{
	"date":     dateValidator,
	"datetime": dateTimeValidator,
	//"password":  passwordValidator,
}

// 自定义验证规则日期：date
func dateValidator(fl validator.FieldLevel) bool {
	date := fl.Field().Interface().(string)
	if date == "" {
		return true
	}
	if _, err := time.Parse(time.DateOnly, date); err != nil {
		return false
	}
	return true
}

// 自定义验证规则日期：datetime
func dateTimeValidator(fl validator.FieldLevel) bool {
	datetime := fl.Field().Interface().(string)
	if datetime == "" {
		return true
	}
	if _, err := time.Parse(time.DateTime, datetime); err != nil {
		return false
	}
	return true
}
