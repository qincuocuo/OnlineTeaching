package mongo

import (
	"context"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"
)

type talk struct {}

var Talk talk

func (talk) Create(ctx context.Context, talk models.Talk) (err error) {
	dbName := talk.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, talk)
	defer span.End()
	err = db.MongoCli.Insert(dbName, talk)
	return
}
