package mongo

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"
)

type exercises struct {
}

var Exercises exercises

func (exercises) Create(ctx context.Context, exercisesDoc models.Exercises) (err error) {
	dbName := exercisesDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, exercisesDoc)
	defer span.End()
	err = db.MongoCli.Insert(dbName, exercisesDoc)
	return
}

func (exercises) FindOne(ctx context.Context, query bson.M) (exercisesDoc models.Exercises, err error) {
	dbName := exercisesDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &exercisesDoc)
	return
}