package chat

import (
	"context"
	"fmt"
	"git.moresec.cn/moresec/go-common/mlog"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"sync"
	"time"
	"webapi/dao/mongo"
	"webapi/models"
)

type Room struct {
	ContentId int

	onlineNum int

	mux     sync.Mutex
	Clients []*ClientId

	// chan
	OnlineChan  chan *Client
	OfflineChan chan *Client
	MessageChan chan *Message

	stopChan chan struct{}
}

func NewChatRoom(courseId int) *Room {
	return &Room{
		ContentId: courseId,

		Clients:     make([]*ClientId, 0),
		OnlineChan:  make(chan *Client, 10),
		OfflineChan: make(chan *Client, 10),
		MessageChan: make(chan *Message, 10),
	}
}

func (r *Room) OnlineNum() int {
	r.mux.Lock()
	defer r.mux.Unlock()

	return r.onlineNum
}

func (r *Room) Process(ctx context.Context, conn *websocket.Conn, userId string) {
	var user models.User
	user, err := mongo.User.FindByUserId(ctx, userId)
	if err != nil {
		mlog.Error("get user info", zap.String("user id", userId), zap.Error(err))
		return
	}

	client := &Client{
		Type: ClientTypeUser,
		User: &user,
		Conn: conn,
	}

	r.mux.Lock()
	r.Clients = append(r.Clients, &ClientId{
		Id:     userId,
		Client: client,
	})
	r.onlineNum++
	r.mux.Unlock()

	r.OnlineChan <- client
	mlog.Info("user online", zap.String("user id", userId), zap.String("user name", user.UserName), zap.Int("online", r.OnlineNum()), zap.Int("room", r.ContentId))

	defer func() {
		r.OfflineChan <- client
		mlog.Info("user offline", zap.String("user id", userId), zap.String("user name", user.UserName), zap.Int("online", r.OnlineNum()), zap.Int("room", r.ContentId))

		r.mux.Lock()
		for i, client := range r.Clients {
			if client.Id == userId {
				r.Clients = append(r.Clients[:i], r.Clients[i+1:]...)
				break
			}
		}
		r.onlineNum--

		r.mux.Unlock()

		if r.OnlineNum() == 0 {
			r.stopChan <- struct{}{}
			release(r)
			mlog.Info("chat room closed", zap.Int("room", r.ContentId))
		}
		conn.Close()
	}()

	for {
		var msg SendMessage
		err = conn.ReadJSON(&msg)
		if err != nil {
			mlog.Info("read message", zap.String("user id", userId), zap.Error(err))
			break
		}

		r.MessageChan <- &Message{
			Type:   MsgCommon,
			Msg:    msg.Msg,
			Sender: userId,
			//Avatar:   user.Avatar,
			Name:     client.User.UserName,
			SendTime: time.Now().Unix(),
		}
	}
}

func (r *Room) Broadcast() {
	for {
		select {
		case client := <-r.OnlineChan:
			msg := &Message{
				Name:     client.User.UserName,
				Msg:      fmt.Sprintf("%s加入聊天室", client.User.UserName),
				Type:     MsgOnline,
				SendTime: time.Now().Unix(),
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("welcome", client.User.UserName)
			r.MessageChan <- msg
		case client := <-r.OfflineChan:
			msg := &Message{
				Name:     client.User.UserName,
				Msg:      fmt.Sprintf("%s离开聊天室", client.User.UserName),
				Type:     MsgOffline,
				SendTime: time.Now().Unix(),
			}

			r.mux.Lock()
			msg.OnlineNum = r.onlineNum
			r.mux.Unlock()

			fmt.Println("bye", client.User.UserName)
			r.MessageChan <- msg
		case msg := <-r.MessageChan:
			fmt.Println("broadcast", msg.Name)
			r.SendMessage(msg)
		case <-time.After(time.Second * 60 * 2):
			mlog.Info("online number", zap.Int("room id", r.ContentId), zap.Int("online", r.OnlineNum()))
		case <-r.stopChan:
			return
		}
	}
}

func (r *Room) SendMessage(msg *Message) {
	switch msg.Type {
	case MsgOnline, MsgOffline:
		for _, clientId := range r.Clients {
			//if clientId.Client.User.Name != msg.Name {
			err := clientId.Client.Conn.WriteJSON(msg)
			if err != nil {
				mlog.Info("send message", zap.Error(err))
			}
		}
		//}
	case MsgRobot:

	case MsgCommon:
		for _, clientId := range r.Clients {
			err := clientId.Client.Conn.WriteJSON(msg)
			if err != nil {
				mlog.Error("send message", zap.Error(err))
			}
			record := models.TalkRecord{
				ContentId: r.ContentId,
				UserId:    clientId.Client.User.UserId,
				Text:      msg.Msg,
				CreateTm:  time.Now(),
			}

			err = mongo.TalkRecord.Create(context.TODO(), record)
			if err != nil {
				mlog.Error("save talk record", zap.Int("content id", r.ContentId), zap.String("user id", record.UserId), zap.Error(err))
			}
		}
	}
}
