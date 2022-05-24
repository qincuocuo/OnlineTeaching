package mongo

import (
	"context"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"

	"github.com/globalsign/mgo/bson"
)

type systemConfig struct{}

var SystemConfig systemConfig

func (systemConfig) Set(ctx context.Context, systemDoc models.SystemCfg) (err error) {
	dbName := systemDoc.CollectName()
	var isExist int
	span, _ := tracking.DbTracking(ctx, dbName, systemDoc)
	defer span.End()
	isExist, err = db.MongoCli.FindCount(dbName, bson.M{})
	if isExist > 0 {
		err = db.MongoCli.Update(dbName, bson.M{}, systemDoc, false)
	} else {
		err = db.MongoCli.Insert(dbName, systemDoc)
	}
	return
}
func (systemConfig) Get(ctx context.Context) (systemDoc models.SystemCfg) {
	var ok bool
	dbName := systemDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, bson.M{})
	defer span.End()
	ok, _ = db.MongoCli.FindOne(dbName, bson.M{}, &systemDoc)
	if !ok {
		systemDoc = models.SystemCfg{
			PwdFlushCycle:     90, //默认密码过期时间(天)
			NoOpExitTm:        30, //默认页面无操作自动退出时间(分钟)
			LoginFailedCount:  5,  //默认允许登录失败次数
			LoginFailedLockTm: 5,  //默认登录失败锁定时间(分钟)
		}
	}
	return
}
