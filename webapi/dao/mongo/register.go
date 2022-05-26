package mongo

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"
)

type register struct {}

var Register register

func (register) Create(ctx context.Context, register models.Register) (err error) {
	dbName := register.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, register)
	defer span.End()
	err = db.MongoCli.Insert(dbName, register)
	return
}

func (register) Update(ctx context.Context, query bson.M, updateSet bson.M) (err error) {
	dbName := (&models.Register{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName)
	defer span.End()
	err = db.MongoCli.Update(dbName, query, updateSet, false)
	return
}

func (register) FindOne(ctx context.Context, query bson.M) (contentDoc models.Register, err error) {
	dbName := contentDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &contentDoc)
	return
}

