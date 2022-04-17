package db

import (
	"fmt"
	"webapi/config"

	"git.moresec.cn/moresec/go-common/mdb"
)

var MongoCli mdb.DBAdaptor

func MongoInit(conf config.MongodbCfg) error {
	var mongoURL = fmt.Sprintf("mongodb://%s:%s@%s:%d/%s", conf.User, conf.Passwd, conf.Host, conf.Port, conf.DbName)
	MongoCli = mdb.NewMongoSession()
	err := MongoCli.Connect(mongoURL)
	if err != nil {
		return err
	}
	return nil
}
