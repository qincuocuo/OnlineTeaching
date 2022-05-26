package preset

import (
	"fmt"
	"log"
	"net/http"
	"time"
	"webapi/config"
	"webapi/middleware/basic"
	"webapi/support"

	"github.com/iris-contrib/swagger/v12"
	"github.com/iris-contrib/swagger/v12/swaggerFiles"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	cover "github.com/kataras/iris/v12/middleware/recover"
	"gopkg.in/natefinch/lumberjack.v2"
)

// PreSetting 注册中间件、定义错误处理
func PreSetting(app *iris.Application) {

	rlog := log.New(&lumberjack.Logger{
		Filename:   config.IrisConf.Log.LogPath + "/request.log",
		MaxSize:    config.IrisConf.Log.MaxSize, // megabytes
		MaxBackups: 3,
		MaxAge:     config.IrisConf.Log.MaxAge, //days
		Compress:   true,                       // disabled by default
	}, "", 3)

	app.Logger().SetLevel(config.IrisConf.JWT.JwtLogLevel)
	c := logger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := fmt.Sprintf("%s | %v | %4v | %s | %s | %s | %v", now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, headerMessage)
		rlog.Println(output)
	}

	customLogger := logger.New(c)
	app.Use(
		customLogger,
		cover.New(),
		//middleware.ServeHTTP
	)

	// ---------------------- 定义错误处理 ------------------------
	app.OnErrorCode(iris.StatusNotFound, customLogger, func(ctx iris.Context) {
		support.SendApiErrorResponse(ctx, support.NotFound, iris.StatusNotFound)
	})
	//app.OnErrorCode(iris.StatusForbidden, customLogger, func(ctx iris.Context) {
	//	ctx.JSON(utils.Error(iris.StatusForbidden, "权限不足", nil))
	//})
	//捕获所有http错误:
	//app.OnAnyErrorCode(customLogger, func(ctx iris.Context) {
	//	//这应该被添加到日志中，因为`logger.Config＃MessageContextKey`
	//	ctx.Values().Set("logger_message", "a dynamic message passed to the logs")
	//	ctx.JSON(utils.Error(500, "服务器内部错误", nil))
	//})

	//删除uri最后的一个 / 不然会触发iris框架的301
	app.WrapRouter(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		path := r.URL.Path
		if len(path) > 1 && path[len(path)-1] == '/' && path[len(path)-2] != '/' {
			path = path[:len(path)-1]
			r.RequestURI = path
			r.URL.Path = path
		}
		next(w, r)
	})

	basic.RegisterIgnoreURLs(config.IrisConf.Web.IgnoreUrls)

	// --------------------------
	// register swagger router
	conf := &swagger.Config{
		//URL: "http://" + config.IrisConf.Web.WebHost + "/swagger/doc.json", //The url pointing to API definition
		URL: "/swagger/doc.json", //The url pointing to API definition
	}
	// use swagger middleware to
	app.Get("/swagger/{any:path}", swagger.CustomWrapHandler(conf, swaggerFiles.Handler))

}
