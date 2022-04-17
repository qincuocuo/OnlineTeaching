package models

import "github.com/globalsign/mgo/bson"

type SystemCfg struct {
	PwdFlushCycle     int `bson:"pwd_flush_cycle"`      //密码更新间隔
	NoOpExitTm        int `bson:"no_oper_exit_tm"`      //无操作退出时间
	LoginFailedCount  int `bson:"login_failed_count"`   //允许登录失败次数
	LoginFailedLockTm int `bson:"login_failed_lock_tm"` //登录失败锁定时间
}

func (systemCfg SystemCfg) CollectName() string {
	return "tb_system_config"
}

type SystemInfo struct {
	Id           bson.ObjectId `bson:"_id,omitempty"`
	WebIP        string        `bson:"web_ip"`         // web服务IP
	WebAddr      string        `bson:"web_addr"`       // https://127.0.0.1:8081 这种
	SystemInitTm int64         `bson:"system_init_tm"` // 系统创建时间
}

func (s SystemInfo) CollectName() string {
	return "tb_system_info"
}
