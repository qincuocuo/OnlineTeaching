package main

import (
	"flag"
	"github.com/pkg/errors"
	"log"
	"webapi/config"
	"webapi/internal/cache"
	"webapi/internal/db"
	"webapi/middleware"
	"webapi/middleware/preset"
	"webapi/middleware/ratelimiter"
	"webapi/middleware/tracking"
	"webapi/router"

	"github.com/kataras/iris/v12"
)

func newApp() *iris.Application {
	app := iris.New()
	preset.PreSetting(app)
	middleware.InitMiddleware(app)
	router.InitRouters(app)
	return app
}

func main() {
	if err := InitService(); err != nil {
		log.Fatal("web_api init service failed")
	}
	app := newApp()
	err := app.Run(iris.Addr(config.IrisConf.Web.WebHost), iris.WithRemoteAddrHeader("X-Forwarded-For"))
	if err != nil {
		log.Fatalf("app run failed :%v\n", err)
	}
}

func InitService() error {
	var err error
	var strPath string
	flag.StringVar(&strPath, "conf_path", "../config/config.yaml", "--conf_path")
	flag.Parse()
	if err = config.CfgInit(strPath); err != nil {
		return errors.WithStack(err)
	}
	// 初始化MongoDB
	if err = db.MongoInit(config.IrisConf.Mongodb); err != nil {
		return errors.WithStack(err)
	}
	// 初始化Redis
	if err = cache.RedisInitPool(config.IrisConf.Redis); err != nil {
		return errors.WithStack(err)
	}

	//初始化SysToken，Jwt防重放
	//if err = jwts.InitSysToken(); err != nil {
	//	return errors.WithStack(err)
	//}
	//初始化链路追踪
	tracking.InitTracking(config.IrisConf.Tracking)
	//初始化日志记录
	config.LogInit()
	// 限流
	ratelimiter.InitRateLimiter()
	return nil
}
