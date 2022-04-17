package mlog

import (
	"go.uber.org/zap"
)

//使用string是为了减少使用Spintf
type ILog interface {
	//new
	//普通日志,如果有args，需要格式化
	Debug(string, ...zap.Field)
	Info(string, ...zap.Field)
	Warn(string, ...zap.Field)
	Error(string, ...zap.Field)
	Panic(string, ...zap.Field)
	Fatal(string, ...zap.Field)
	//需要格式化日志 ，最后一个是context
	Debugf(string, ...interface{})
	Infof(string, ...interface{})
	Warnf(string, ...interface{})
	Errorf(string, ...interface{})
	Panicf(string, ...interface{})
	Fatalf(string, ...interface{})
}
