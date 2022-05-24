package mongo

import (
	"context"
	"github.com/globalsign/mgo/bson"
	"webapi/internal/db"
	"webapi/internal/utils"
	"webapi/middleware/tracking"
	"webapi/models"
)

type content struct{}

var Content content

func (content) FindCount(ctx context.Context, query bson.M) (count int, err error) {
	dbName := (&models.LearningContent{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	count, err = db.MongoCli.FindCount(dbName, query)
	return
}

func (content) FindOne(ctx context.Context, query bson.M) (contentDoc models.LearningContent, err error) {
	dbName := contentDoc.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	_, err = db.MongoCli.FindOne(dbName, query, &contentDoc)
	return
}

func (content) FindSortByLimitAndSkip(ctx context.Context, query bson.M, page int, pageSize int, sorter string) (courseDoc []models.LearningContent, err error) {
	dbName := (&models.LearningContent{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, query)
	defer span.End()
	limit := pageSize
	skip := utils.GetPageStart(page, pageSize)
	err = db.MongoCli.FindSortByLimitAndSkip(dbName, query, courseDoc, limit, skip, sorter)
	return
}

func (content) GetMaxId(ctx context.Context) (uid int) {
	var courseDoc []models.LearningContent
	dbName := (&models.LearningContent{}).CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, bson.M{})
	defer span.End()
	_ = db.MongoCli.FindSortByLimitAndSkip(dbName, bson.M{}, &courseDoc, 1, 0, "-uid")
	uid = courseDoc[0].ContentId + 1
	return
}

