package chat

import (
	"context"
	"git.moresec.cn/moresec/go-common/mlog"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"sync"
	"time"
)

var (
	Pool  sync.Pool
	Rooms = make(map[int]*Room)
)

func init() {
	Pool = sync.Pool{
		New: func() interface{} {
			return &Room{
				Clients: make([]*ClientId, 0),

				OnlineChan:  make(chan *Client, 10),
				OfflineChan: make(chan *Client, 10),
				MessageChan: make(chan *Message, 10),
			}
		},
	}

	go func() {
		for {
			select {
			case <-time.After(time.Second * 60 * 5):
				mlog.Info("chat rooms", zap.Any("chat rooms", Rooms))
			}
		}
	}()
}

func Process(ctx context.Context, courseId int, userId string, conn *websocket.Conn) {
	room, ok := Rooms[courseId]
	if !ok {
		room = Pool.Get().(*Room)
		room.CourseId = courseId
		Rooms[courseId] = room

		mlog.Info("New room created", zap.Int("room id", courseId), zap.Any("room", room))

		go room.Broadcast()
	}
	room.Process(ctx, conn, userId)
}

func release(room *Room) {
	Pool.Put(room)
	delete(Rooms, room.CourseId)
}
