package tracking

import (
	"context"
	"fmt"
	"webapi/utils"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"

	"git.moresec.cn/moresec/go-common/mlog"
	irisCtx "github.com/kataras/iris/v12/context"
	ztrace "github.com/tal-tech/go-zero/core/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

const TracingName = "sun_tracing"

//
//var (
//	UserIdKey = attribute.Key("uid")
//)

func InitTracking(trackingAddress string) {
	if trackingAddress == "" {
		return
	}

	ztrace.StartAgent(ztrace.Config{
		Name:     "webapi_srv",
		Endpoint: trackingAddress,
		Batcher:  "jaeger",
		Sampler:  1.0,
	})
}

func SetIrisContext(irisCtx irisCtx.Context, commonCtx context.Context) {
	irisCtx.Values().Set(TracingName, commonCtx)
}

func StartSpan(ctx context.Context, spanName string, opts ...trace.SpanStartOption) (trace.Span, context.Context) {
	ctx, span := otel.Tracer(ztrace.TraceName).Start(ctx, spanName, opts...)
	return span, ctx
}

func GetIrisContext(irisCtx irisCtx.Context) (commonCtx context.Context) {
	var ok bool
	commonCtx, ok = irisCtx.Values().Get(TracingName).(context.Context)
	if !ok || commonCtx == nil {
		mlog.Warn("ctx is nil ! ")
		return context.Background()
	}
	return
}

func SetIrisSpan(irisCtx irisCtx.Context, span trace.Span) {
	irisCtx.Values().Set("Span", span)
}

func GetIrisSpan(irisCtx irisCtx.Context) (span trace.Span) {
	span, ok := irisCtx.Values().Get("Span").(trace.Span)
	if !ok {
		return nil
	}
	return span
}

func DbTracking(commonCtx context.Context, dbName string, query ...interface{}) (trace.Span, context.Context) {
	funcName := utils.Invoke.GetFuncName(2)
	span, ctx := StartSpan(
		commonCtx,
		fmt.Sprintf("(%s)%s", dbName, funcName),
		trace.WithSpanKind(trace.SpanKindClient),
		trace.WithAttributes(
			semconv.DBSystemMongoDB,
			// semconv.DBOperationKey.String(evt.CommandName),
			semconv.DBNameKey.String(dbName),
			//semconv.NetPeerNameKey.String(hostname),
			// semconv.NetPeerPortKey.Int(port),
			semconv.NetTransportTCP,
		),
	)
	return span, ctx
}

func RedisTracking(commonCtx context.Context, rdxKey string, params ...interface{}) (trace.Span, context.Context) {
	funcName := utils.Invoke.GetFuncName(2)
	span, ctx := StartSpan(commonCtx, fmt.Sprintf("(%s)%s", rdxKey, funcName))
	// span.SetTag("func_name", funcName)
	// span.SetTag("rdx_key", rdxKey)
	// span.SetTag("other_params", params)
	return span, ctx
}
