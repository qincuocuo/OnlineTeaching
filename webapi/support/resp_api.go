package support

import (
	"fmt"
	"net/http"

	"github.com/globalsign/mgo/bson"

	"github.com/kataras/iris/v12"

	"git.moresec.cn/moresec/go-common/mlog"
	"github.com/kataras/iris/v12/context"
	"go.uber.org/zap"
)

// makeResponse 创建http响应
func makeResponse(ctx context.Context, data interface{}, msg string, errMsg string, code int) {
	var spanId string
	resp := bson.M{
		"code":       code,
		"error":      errMsg,
		"data":       data,
		"message":    msg,
		"request_id": fmt.Sprintf("%s", spanId),
	}
	_, err := ctx.JSON(resp)
	if err != nil {
		mlog.Info("make response error ", zap.Error(err))
	}
}

// SendApiResponse 构建正确的接口返回
func SendApiResponse(ctx context.Context, data interface{}, msg string) {
	if msg == "" {
		msg = "success"
	}
	makeResponse(ctx, data, msg, "", iris.StatusOK)
}

// SendApiErrorResponse 构建错误的接口返回
func SendApiErrorResponse(ctx context.Context, msg string, statusCode int) {
	makeResponse(ctx, nil, msg, msg, statusCode)
}

// SetAuthCookie 设置登录cookie
func SetAuthCookie(ctx context.Context, token string) {
	if token != "" {
		ctx.SetCookie(&http.Cookie{
			Name:     Auth,
			Value:    token,
			Path:     "/",
			HttpOnly: true,
		})
	}
}
