package main

import (
	"log"
	"webapi"
	"webapi/config"
	_ "webapi/docs"
	"webapi/middleware"
	"webapi/middleware/preset"
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
	if err := webapi.InitService(); err != nil {
		log.Fatal("web_api init service failed")
	}
	app := newApp()
	err := app.Run(iris.Addr(config.IrisConf.Web.WebHost), iris.WithRemoteAddrHeader("X-Forwarded-For"))
	if err != nil {
		log.Fatalf("app run failed :%v\n", err)
	}
}
