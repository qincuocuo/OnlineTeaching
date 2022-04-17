package mongo

import (
	"common/models"
	"context"
	"webapi/internal/db"
	"webapi/middleware/tracking"

	"github.com/globalsign/mgo/bson"
)

type user struct{}

var User user

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

func (user) FindByNameRole(ctx context.Context, name string, role int) (userDoc models.User, err error) {
	dbName := userDoc.CollectName()
	query := bson.M{"username": name, "role": role}
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &userDoc)
	return
}

func (user) GetMaxUid(ctx context.Context) (uid int) {
	var userDoc []models.User
	dbName := (&models.User{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, bson.M{})
	defer span.End()
	_ = db.MongoCli.FindSortByLimitAndSkip(dbName, bson.M{}, &userDoc, 1, 0, "-uid")
	uid = userDoc[0].UID + 1
	return
}
