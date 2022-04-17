package ratelimiter

import (
	"webapi/config"

	"go.uber.org/ratelimit"
)

var RateLimit ratelimit.Limiter

func InitRateLimiter() {
	if config.IrisConf.Web.RateLimit != 0 {
		RateLimit = ratelimit.New(int(config.IrisConf.Web.RateLimit))
	}
}
