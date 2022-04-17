/*
@Time : 2020-10-27 10:39
@Author : gaoxl@moresec.cn
@Description:
@Software: GoLand
*/
package mlog

import (
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"

	"go.uber.org/zap/zapcore"
)

type Config struct {
	Path          string `yaml:"path" json:"path"`               //  日志输出目录
	ProjectName   string `yaml:"projectName" json:"projectName"` // 日志文件名称
	Level         string `yaml:"level" json:"level"`             // 日志初始等级
	AddCaller     bool   `yaml:"addCaller" json:"addCaller"`     // 是否添加调用者信息
	Prefix        string `yaml:"prefix" json:"prefix"`           // 日志前缀
	MaxSize       int    `yaml:"maxSize" json:"maxSize"`         // 日志输出文件最大长度，超过改值则截断
	MaxAge        int    `yaml:"maxAge" json:"maxAge"`           // 日志保存周期
	MaxBackup     int    `yaml:"maxBackup" json:"maxBackup"`     //日志备份
	CallerSkip    int    `yaml:"callerSkip" json:"callerSkip"`
	Async         bool   `yaml:"async" json:"async"`
	Queue         bool   `yaml:"queue" json:"queue"`
	Debug         bool   `yaml:"debug" json:"debug"`
	Interval      int64  `yaml:"interval" json:"interval"` // 日志磁盘刷盘间隔
	QueueSleep    int64
	Core          zapcore.Core
	EncoderConfig *zapcore.EncoderConfig
	Fields        []zap.Field // 日志初始化字段
}

func NewConfigFromJsonData(data string) *Config {
	var conf Config
	if err := jsoniter.UnmarshalFromString(data, &conf); err != nil {
		panic(err)
	}
	return &conf
}

func (config Config) Build(projectName string) *Logger {
	config.ProjectName = projectName
	if config.EncoderConfig == nil {
		config.EncoderConfig = DefaultZapConfig()
	}

	if config.Debug {
		config.EncoderConfig.EncodeLevel = DebugEncodeLevel
	}
	logger := newLogger(&config)
	return logger
}
