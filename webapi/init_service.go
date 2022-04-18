package webapi

import (
	"flag"
	"webapi/config"
	"webapi/internal/cache"
	"webapi/internal/db"
	"webapi/middleware/ratelimiter"
	"webapi/middleware/tracking"

	"github.com/pkg/errors"
)

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
