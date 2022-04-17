package middleware

import (
	"fmt"
	"webapi/config"
	"webapi/middleware/basic"
	"webapi/middleware/jwts"
	"webapi/middleware/ratelimiter"
	"webapi/middleware/tracking"

	semconv "go.opentelemetry.io/otel/semconv/v1.7.0"
	"go.opentelemetry.io/otel/trace"

	"github.com/kataras/iris/v12/context"

	"github.com/kataras/iris/v12"
)

// InitMiddleware 初始化中间件
func InitMiddleware(app *iris.Application) {
	app.Use(rateLimiterMiddleware) //限流中间件
	app.Use(tracingMiddleware)     //链路追踪中间件
	app.Use(jwtMiddleware)         //Jwt登录中间件
}

// 限流中间件
func rateLimiterMiddleware(ctx context.Context) {
	if config.IrisConf.Web.RateLimit != -1 && config.IrisConf.Web.RateLimit != 0 {
		ratelimiter.RateLimit.Take()
	}
	ctx.Next()
}

// 链路追踪中间件
func tracingMiddleware(ctx context.Context) {
	span, trackCtx := tracking.StartSpan(
		ctx.Request().Context(),
		fmt.Sprintf("%s[%s]", ctx.Request().RequestURI, ctx.Method()),
		trace.WithAttributes(semconv.NetAttributesFromHTTPRequest("tcp", ctx.Request())...),
		trace.WithAttributes(semconv.EndUserAttributesFromHTTPRequest(ctx.Request())...),
		trace.WithAttributes(semconv.HTTPServerAttributesFromHTTPRequest("webapi_srv", ctx.Path(), ctx.Request())...),
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()

	tracking.SetIrisContext(ctx, trackCtx)
	tracking.SetIrisSpan(ctx, span)

	ctx.Next()
}

// JwtMiddleware Jwt登录中间件
func jwtMiddleware(ctx context.Context) {
	if !basic.CheckURL(ctx.Path()) {
		if ctx.Values().Get(jwts.DefaultContextKey) == nil {
			//jwt token拦截
			if !jwts.Serve(ctx) {
				ctx.StatusCode(iris.StatusUnauthorized)
				ctx.StopExecution()
			}
		}
	}
	ctx.Next()
}
