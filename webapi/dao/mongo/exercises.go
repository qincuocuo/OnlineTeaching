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

func (exercises) Update(ctx context.Context, query, upset bson.M) (err error) {
	dbName := (&models.Exercises{}).CollectName()

	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()

	err = db.MongoCli.Update(dbName, query, upset, false)

	return
}

func (exercises) IsExist(ctx context.Context, query bson.M) bool {
	var exercisesDoc []models.Exercises
	dbName := (&models.Exercises{}).CollectName()

	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()

	if err := db.MongoCli.FindAll(dbName, query, &exercisesDoc); err != nil || len(exercisesDoc) == 0 {
		return false
	}

	return true
}
