package mongo

import (
	"context"
	"webapi/internal/db"
	"webapi/middleware/tracking"
	"webapi/models"
)

type talk struct{}
type talkRecord struct{}

var (
	Talk       talk
	TalkRecord talkRecord
)

func (talk) Create(ctx context.Context, talk models.Talk) (err error) {
	dbName := talk.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, talk)
	defer span.End()
	err = db.MongoCli.Insert(dbName, talk)
	return
}

func (talkRecord) Create(ctx context.Context, talkRecord models.TalkRecord) (err error) {
	dbName := talkRecord.CollectName()
	span, _ := tracking.DbTracking(ctx, dbName, talkRecord)
	defer span.End()
	err = db.MongoCli.Insert(dbName, talkRecord)
	return
}
