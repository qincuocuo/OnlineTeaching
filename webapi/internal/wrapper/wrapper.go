package wrapper

import (
	"encoding/json"
	"fmt"
	"runtime/debug"
	"webapi/support"

	"reflect"
	"strings"

	"gopkg.in/go-playground/validator.v9"

	"go.uber.org/zap"

	"github.com/kataras/iris/v12"

	"git.moresec.cn/moresec/go-common/mlog"
)

type (
	ApiHandler func(ctx *Context, reqBody interface{}) error
)

type ApiConfig struct {
	ReqType support.CheckType
}

var validate = validator.New()

// ApiWrapper params 约定 Api封装器
func ApiWrapper(ctx *Context, handler ApiHandler, paramChecker bool, reqBody interface{}, params ...interface{}) {
	defer func() {
		if r := recover(); r != nil {
			debug.PrintStack()
			errStr := fmt.Sprintf("recover:%v", r)
			mlog.Error("recover error", zap.String("msg", errStr))
			support.SendApiErrorResponse(ctx, "recover error", iris.StatusInternalServerError)
		}
	}()
	var err error
	var paramErr map[string]string
	//请求Body校验
	if reqBody != nil {
		if len(params) == 0 { //参数为空
			mlog.Error("ApiWrapper传入params为空")
			support.SendApiErrorResponse(ctx, "params is empty", iris.StatusInternalServerError)
			return
		}
		//请求类型校验 form|json
		config := params[0].(*ApiConfig)
		switch config.ReqType {
		case support.CHECKTYPE_FORM:
			err = ctx.ReadForm(reqBody)
		case support.CHECKTYPE_JSON:
			err = ctx.ReadJSON(reqBody)
		}
		if err != nil && !iris.IsErrPath(err) {
			mlog.Error("Api Wrapper reqBody parse failed", zap.Error(err))
			support.SendApiErrorResponse(ctx, "parse reqBody failed", iris.StatusInternalServerError)
			return
		}
		//请求参数校验
		if paramChecker {
			if err, paramErr = checkParam(config.ReqType, reqBody); err != nil || paramErr != nil {
				if err != nil {
					mlog.Error("checker param error", zap.Error(err))
					support.SendApiErrorResponse(ctx, "checker param failed failed", iris.StatusOK)
				} else {
					if msg, err := json.Marshal(paramErr); err == nil {
						support.SendApiErrorResponse(ctx, string(msg), iris.StatusOK)
					} else {
						mlog.Error("marshal err param msg failed", zap.Error(err))
						support.SendApiErrorResponse(ctx, "checker param failed failed", iris.StatusInternalServerError)
					}
				}
				return
			}
		}
	}
	if err := handler(ctx, reqBody); err != nil {
		mlog.Error("handler exec failed", zap.Error(err))
		support.SendApiErrorResponse(ctx, "handler exec failed", iris.StatusInternalServerError)
	}
}

// 参数校验函数
func checkParam(reqType support.CheckType, reqBody interface{}) (error, map[string]string) {
	if err := validate.Struct(reqBody); err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			return err, nil
		}
		paramErr := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			st := reflect.TypeOf(reqBody)
			if param, ok := st.Elem().FieldByName(strings.Split(err.StructField(), "[")[0]); ok {
				if reqType == support.CHECKTYPE_JSON {
					tag := param.Tag.Get("json")
					paramErr[tag] = "Invalid input"
				} else if reqType == support.CHECKTYPE_FORM {
					tag := param.Tag.Get("form")
					paramErr[tag] = "Invalid input"
				}
			}
		}
		return nil, paramErr
	}
	return nil, nil
}
