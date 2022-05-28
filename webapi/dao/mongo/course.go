package mongo

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"webapi/internal/db"
	"webapi/internal/utils"
	"webapi/middleware/tracking"
	"webapi/models"
)

type course struct{}

var Course course

func (course) FindOne(ctx context.Context, query bson.M) (courseDoc models.Course, err error) {
	dbName := courseDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &courseDoc)
	return
}

func (course) FindAll(ctx context.Context, query bson.M) (courseDoc []models.Course, err error) {
	dbName := (&models.Course{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	err = db.MongoCli.FindAll(dbName, query, &courseDoc)
	return
}

func (course) Create(ctx context.Context, course models.Course) (err error) {
	dbName := course.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, course)
	defer span.End()
	err = db.MongoCli.Insert(dbName, course)
	return
}

func (course) Update(ctx context.Context, query bson.M, updateSet bson.M) (err error) {
	dbName := (&models.Course{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName)
	defer span.End()
	err = db.MongoCli.Update(dbName, query, updateSet, false)
	return
}

func (course) Delete(ctx context.Context, query bson.M) (err error) {
	dbName := (&models.Course{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName)
	defer span.End()
	err = db.MongoCli.RemoveAll(dbName, query)
	return
}

func (course) FindSortByLimitAndSkip(ctx context.Context, query bson.M, page int, pageSize int, sorter string) (courseDoc []models.Course, err error) {
	dbName := (&models.Course{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	limit := pageSize
	skip := utils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindSortByLimitAndSkip(dbName, query, courseDoc, limit, skip, sorter)
	return
}

func (course) GetMaxId(ctx context.Context) (uid int) {
	var courseDoc []models.Course
	dbName := (&models.Course{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, bson.M{})
	defer span.End()
	_ = db.MongoCli.FindSortByLimitAndSkip(dbName, bson.M{}, &courseDoc, 1, 0, "-course_id")
	uid = courseDoc[0].CourseId + 1
	return
}
