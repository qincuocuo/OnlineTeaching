package mlog

import (
	"git.moresec.cn/moresec/go-common/mlog/zaplog"
	"go.uber.org/zap"
)

//默认
var l ILog = zaplog.New()

//设置
func SetLogger(ll ILog) {
	l = ll
}

func GetLogger() ILog {
	return l
}

//普通日志
func Debug(msg string, filed ...zap.Field) {
	l.Debug(msg, filed...)
}
func Info(msg string, filed ...zap.Field) {
	l.Info(msg, filed...)
}
func Warn(msg string, filed ...zap.Field) {
	l.Warn(msg, filed...)
}
func Error(msg string, filed ...zap.Field) {
	l.Error(msg, filed...)
}
func Panic(msg string, filed ...zap.Field) {
	l.Panic(msg, filed...)
}
func Fatal(msg string, filed ...zap.Field) {
	l.Fatal(msg, filed...)
}

//其他日志 如：HTTP RPC日志
func Debugf(format string, args ...interface{}) {
	l.Debugf(format, args...)
}
func Infof(format string, args ...interface{}) {
	l.Infof(format, args...)
}
func Warnf(format string, args ...interface{}) {
	l.Warnf(format, args...)
}
func Errorf(format string, args ...interface{}) {
	l.Errorf(format, args...)
}
func Panicf(format string, args ...interface{}) {
	l.Panicf(format, args...)
}
func Fatalf(format string, args ...interface{}) {
	l.Fatalf(format, args...)
}
