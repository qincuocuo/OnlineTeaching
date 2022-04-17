package mlog

import (
	"context"
	"testing"

	"git.moresec.cn/moresec/go-common/mlog/conf"
	"git.moresec.cn/moresec/go-common/mlog/zaplog"

	"go.uber.org/zap"
)

//Debug(string, ...zap.Field)
//Info(string, ...zap.Field)
//Warn(string, ...zap.Field)
//Error(string, ...zap.Field)
//Panic(string, ...zap.Field)
//Fatal(string, ...zap.Field)
////需要格式化日志 ，最后一个是context
//Debugf(string, ...interface{})
//Infof(string, ...interface{})
//Warnf(string, ...interface{})
//Errorf(string, ...interface{})
//Panicf(string, ...interface{})
//Fatalf(string, ...interface{})

func TestSetLogger(t *testing.T) {
	//设置为当前目录下 //设置级别
	SetLogger(zaplog.New(
		conf.WithProjectName("zap test"),
		conf.WithLogPath("tmp"),
		conf.WithLogLevel("debug"),
	))
	Debug("this is zap")
	Warn("this is warn")
	//Fatal("this is fatal")
	//Error("this is error.")
	//Panic("this is panic.")
	Debug("hello", zap.String("key", "world"))
	Errorf("failed to open file:%s", "ds_server", context.Background())
	Infof("hello %s %s", "world", "world", context.Background())
}

func TestNewLogger(t *testing.T) {
	log := Config{
		Debug:     true,
		AddCaller: true,
		Level:     "warn",
	}.Build("test1")
	log.Info("hello", String("key", "world"), Int("test", 1))
	log.With(String("mod", "test")).Info("hello", String("key", "world"), Int("test", 1))
	log.With(String("mod", "kv")).Infow("hello", "key", "world", "test", 1)
	log.With(String("mod", "fmt")).Infof("hello %s", "world")
	log.With(String("mod", "test")).Debug("hello", String("key", "world"), Int("test", 1))
	log.With(String("mod", "test")).Warn("hello", String("key", "world"), Int("test", 1))
	log.With(String("mod", "test")).Error("hello", String("key", "world"), Int("test", 1))
	log.With(String("mod", "test")).Fatal("hello", String("key", "world"), Int("test", 1))

	log.With(String("mod", "test")).Panic("hello", String("key", "world"), Int("test", 1))

}
