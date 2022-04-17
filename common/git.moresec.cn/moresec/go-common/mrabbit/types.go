package mrabbit

import (
	"log"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
)

var (
	ErrSend  = errors.New("send message error")
	ErrRead  = errors.New("read message error")
	ErrRetry = errors.New("beyond retry times")
)

type RBQLogger interface {
	Printf(string, ...interface{})
}

type defaultLogger struct {
}

func (l *defaultLogger) Printf(s string, param ...interface{}) {
	log.Printf(s+"\n", param...)
}

var rbqlog RBQLogger = &defaultLogger{}

func SetLogger(l RBQLogger) {
	rbqlog = l
}

type RabbitMsg struct {
	exchangeName string
	routeKey     string
	body         *amqp.Publishing
}

type ConnectOption struct {
	User     string
	Password string
	Host     string
	Port     int32
	VHost    string
}

type QueueConfig struct {
	QueueName string
	RouteKey  string

	ExchangeName string
	ExchangeType string
	Durable      bool
	Exclusive    bool
}

type RabbitSender interface {
	Write(msg *RabbitMsg) error
	Read() chan *RabbitMsg

	OnError()
	OnClose()
}

type DefaultChannel struct {
	msgqueue chan *RabbitMsg
}

func NewDefaultChannel() *DefaultChannel {
	return &DefaultChannel{
		msgqueue: make(chan *RabbitMsg, 1),
	}
}

func (dc *DefaultChannel) Write(msg *RabbitMsg) error {
	dc.msgqueue <- msg
	return nil
}

func (dc *DefaultChannel) Read() chan *RabbitMsg {
	return dc.msgqueue
}

func (dc *DefaultChannel) OnError() {

}

func (dc *DefaultChannel) OnClose() {

}
