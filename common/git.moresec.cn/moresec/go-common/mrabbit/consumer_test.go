package mrabbit

import (
	"testing"

	"github.com/streadway/amqp"
)

type TestReceiver struct {
	queueName string
	routeKey  string
}

func (r *TestReceiver) QueueName() string {
	return r.queueName
}
func (r *TestReceiver) RouteKey() string {
	return r.routeKey
}
func (r *TestReceiver) OnError(err error) {
	rbqlog.Printf("TestReceiver onError. Err:%s", err.Error())
}
func (r *TestReceiver) OnReceiver(d amqp.Delivery) bool {
	rbqlog.Printf("[Receive] message=%v", string(d.Body))
	return true
}
func (r *TestReceiver) QueueConfig() QueueConfig {
	return QueueConfig{}
}

func TestRabbitMQ_Consumer(t *testing.T) {
	const (
		User   string = "moresec"
		Passwd string = "moresec@drip"
		Host   string = "192.168.30.140"
		Port   int32  = 5672
		VHost  string = "moresec_vhost"
	)
	rbqConsumer := NewRabbitMQ(ConnectOption{
		User:     User,
		Password: Passwd,
		Host:     Host,
		Port:     Port,
		VHost:    VHost,
	}, "drip.exchange", "direct")
	if rbqConsumer == nil {
		t.Fail()
	}

	receiver := &TestReceiver{
		"drip_route_msg",
		"drip_route_msg",
	}
	rbqConsumer.RegisterReceiver(receiver)
	rbqConsumer.Start()
}
