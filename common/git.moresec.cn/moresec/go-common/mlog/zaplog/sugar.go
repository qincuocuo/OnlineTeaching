package zaplog

import (
	"context"
	"fmt"

	"git.moresec.cn/moresec/go-common/mlog/tracer"
	"go.uber.org/zap"
)

//
func (l *Log) Debug(s string, filed ...zap.Field) {
	l.logger.Debug(s, filed...)
}
func (l *Log) Info(s string, filed ...zap.Field) {

	l.logger.Info(s, filed...)
}
func (l *Log) Warn(s string, filed ...zap.Field) {

	l.logger.Warn(s, filed...)
}
func (l *Log) Error(s string, filed ...zap.Field) {

	l.logger.Error(s, filed...)
}
func (l *Log) Panic(s string, filed ...zap.Field) {

	l.logger.Panic(s, filed...)
}
func (l *Log) Fatal(s string, filed ...zap.Field) {

	l.logger.Fatal(s, filed...)
}

//

//判断其他类型--start
func getOtherFileds(format string, args ...interface{}) (string, []zap.Field) {
	//判断是否有context
	l := len(args)
	if l > 0 {
		if ctx, ok := args[l-1].(context.Context); ok {
			return fmt.Sprintf(format, args[:l-1]...), []zap.Field{
				zap.String("trace_id", tracer.GetTraceId(ctx)),
				zap.String("parent_id", tracer.GetParentId(ctx)),
				zap.String("span_id", tracer.GetSpanId(ctx)),
			}
		}

		return fmt.Sprintf(format, args...), []zap.Field{}
	}
	return format, []zap.Field{}
}

func (l *Log) Debugf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Debug(s, f...)
}

func (l *Log) Infof(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Info(s, f...)
}

func (l *Log) Warnf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Warn(s, f...)
}

func (l *Log) Errorf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Error(s, f...)
}

func (l *Log) Panicf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Panic(s, f...)
}

func (l *Log) Fatalf(format string, args ...interface{}) {
	s, f := getOtherFileds(format, args...)
	l.logger.Fatal(s, f...)
}
