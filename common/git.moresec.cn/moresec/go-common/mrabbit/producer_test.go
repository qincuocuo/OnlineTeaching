package mrabbit

import (
	"testing"
	"time"

	"github.com/streadway/amqp"
)

func TestRBQProducer_Func(t *testing.T) {
	const (
		User   string = "moresec"
		Passwd string = "moresec@drip"
		Host   string = "192.168.30.140"
		Port   int32  = 5672
		VHost  string = "moresec_vhost"
	)

	rbqclient := NewRBQProducer(ConnectOption{
		User:     User,
		Password: Passwd,
		Host:     Host,
		Port:     Port,
		VHost:    VHost,
	}, nil)

	var queueList []QueueConfig
	queueList = append(queueList, QueueConfig{
		QueueName:    "hello_3",
		RouteKey:     "hello_3",
		ExchangeName: "test",
		ExchangeType: "direct",
		Durable:      true,
	})
	rbqclient.RegisterChannel(queueList)
	go rbqclient.Start()

	//
	var err error
	for i := 0; i < 10000; i++ {
		err = rbqclient.Send("test", "hello_3", &amqp.Publishing{
			ContentType: "application/json",
			Body:        []byte(`{"test":"testValue"}`),
		})
		if err != nil {
			t.Error("send error.")
		}
		time.Sleep(1 * time.Second)
	}
	rbqclient.Close()
}
