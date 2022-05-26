package config

import (
	"git.moresec.cn/moresec/go-common/mlog/zaplog"
	"gopkg.in/yaml.v2"
	"io/ioutil"

	"git.moresec.cn/moresec/go-common/mlog"
	"git.moresec.cn/moresec/go-common/mlog/conf"
)

var IrisConf Config

func CfgInit(strPath string) (err error) {
	strConf, err := ioutil.ReadFile(strPath)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(strConf, &IrisConf); err != nil {
		return err
	}
	return
}

func LogInit() {
	mlog.SetLogger(zaplog.New(
		conf.WithProjectName("sun_webapi"),
		conf.WithLogPath(IrisConf.Log.LogPath),
		conf.WithLogLevel(IrisConf.Log.LogLevel),
		conf.WithMaxSize(IrisConf.Log.MaxSize),
		conf.WithMaxAge(IrisConf.Log.MaxAge),
		conf.WithIsStdOut(IrisConf.Log.IsStdOut),
		conf.WithLogName("sun_webapi"),
	))
}
