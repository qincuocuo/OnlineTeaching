package mongo

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"
)

type user struct{}

var User user

func (user) FindOne(ctx context.Context, query bson.M) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) UserCountByGradeAndClass(ctx context.Context, query bson.M) (count int, err error) {
	dbName := (&models.User{}).CollectName()

	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()

	return db.MongoCli.FindCount(dbName, query)
}

func (user) IsExist(ctx context.Context, query bson.M) bool {
	var userDoc []models.User
	dbName := (&models.User{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	if err := db.MongoCli.FindAll(dbName, query, &userDoc); err != nil || len(userDoc) == 0 {
		return false
	}
	return true
}

func (user) Get(ctx context.Context, uid int) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	query := bson.M{"uid": uid}
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) Update(ctx context.Context, query bson.M, upset bson.M) (err error) {
	dbName := (&models.User{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	err = db.MongoCli.Update(dbName, query, upset, false)
	return
}

func (user) Create(ctx context.Context, user models.User) (err error) {
	dbName := user.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, user)
	defer span.End()
	err = db.MongoCli.Insert(dbName, user)
	return
}

func (user) GetInfo(ctx context.Context, uid int) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	query := bson.M{"uid": uid}
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) FindByUserId(ctx context.Context, id string) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	query := bson.M{"user_id": id}
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) FindByNameRole(ctx context.Context, name string, role int) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	query := bson.M{"username": name, "role": role}
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}
