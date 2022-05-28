package chat

import (
	"github.com/gorilla/websocket"
	"webapi/models"
)

type MessageType int

const (
	_ MessageType = iota
	MsgOnline
	MsgOffline
	MsgRobot
	MsgCommon
)

const (
	ClientTypeUser = "User"
	ClientTypeBot  = "Bot"
)

type Message struct {
	Type      MessageType `json:"type,omitempty"`
	Msg       string      `json:"msg,omitempty"`
	Name      string      `json:"name,omitempty"`
	Sender    string      `json:"sender,omitempty"`
	Avatar    string      `json:"avatar,omitempty"`
	OnlineNum int         `json:"online_num,omitempty"`
	SendTime  int64       `json:"send_time,omitempty"`
}

type SendMessage struct {
	Msg string `json:"msg" form:"msg"`
}

type ClientId struct {
	Id     string  `json:"id"`
	Client *Client `json:"client"`
}

type Client struct {
	Conn *websocket.Conn
	Type string
	User *models.User
}
