package mlog

import (
	"context"
	"fmt"
	"git.moresec.cn/moresec/go-common/mlog/tracer"
	otel_trace "go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const (
	FieldKeyTraceID = "trace_id"
	FieldKeySpanID  = "span_id"
)

type traceLogger struct {
	ctx     context.Context
	TraceID string `json:"trace_id"`
	SpanID  string `json:"span_id"`
}

func (l *traceLogger) formatTraceFields(fields []zap.Field) []zap.Field {
	var isSkipAddTraceID bool
	var isSkipAddSpanID bool

	for _, field := range fields {
		if field.Key == FieldKeyTraceID {
			isSkipAddTraceID = true
		}

		if field.Key == FieldKeySpanID {
			isSkipAddSpanID = true
		}
	}

	if !isSkipAddTraceID {
		fields = append(fields, zap.String("trace_id", l.TraceID))
	}

	if !isSkipAddSpanID {
		fields = append(fields, zap.String("span_id", l.SpanID))
	}

	return fields
}

func (l *traceLogger) Debug(msg string, field ...zap.Field) {
	Debug(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Info(msg string, field ...zap.Field) {
	Info(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Warn(msg string, field ...zap.Field) {
	Warn(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Error(msg string, field ...zap.Field) {
	Error(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Panic(msg string, field ...zap.Field) {
	Panic(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Fatal(msg string, field ...zap.Field) {
	Fatal(msg, l.formatTraceFields(field)...)
}

func (l *traceLogger) Debugf(format string, v ...interface{}) {
	l.Debug(fmt.Sprintf(format, v...))
}

func (l *traceLogger) Infof(format string, v ...interface{}) {
	l.Info(fmt.Sprintf(format, v...))
}

func (l *traceLogger) Warnf(format string, v ...interface{}) {
	l.Warn(fmt.Sprintf(format, v...))
}

func (l *traceLogger) Errorf(format string, v ...interface{}) {
	l.Error(fmt.Sprintf(format, v...))
}

func (l *traceLogger) Panicf(format string, v ...interface{}) {
	l.Panic(fmt.Sprintf(format, v...))
}

func (l *traceLogger) Fatalf(format string, v ...interface{}) {
	l.Fatal(fmt.Sprintf(format, v...))
}

// WithContext sets ctx to log, for keeping tracing information.
func WithContext(ctx context.Context) ILog {
	return &traceLogger{
		ctx:     ctx,
		TraceID: traceIdFromContext(ctx),
		SpanID:  spanIdFromContext(ctx),
	}
}

func spanIdFromContext(ctx context.Context) string {
	span := otel_trace.SpanFromContext(ctx)
	if span.IsRecording() {
		return span.SpanContext().SpanID().String()
	} else {
		return tracer.GetSpanId(ctx)
	}
}

func traceIdFromContext(ctx context.Context) string {
	span := otel_trace.SpanFromContext(ctx)
	if span.IsRecording() {
		return span.SpanContext().TraceID().String()
	} else {
		return tracer.GetTraceId(ctx)
	}
}
