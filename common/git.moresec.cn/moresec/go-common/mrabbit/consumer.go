package mrabbit

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/streadway/amqp"
)

type Receiver interface {
	OnError(error)
	OnReceiver(delivery amqp.Delivery) bool
	QueueConfig() QueueConfig
}

type RabbitMQ struct {
	wg sync.WaitGroup

	opt ConnectOption

	mu   sync.Mutex
	conn *amqp.Connection

	exchangeName string
	exchangeType string
	receivers    []Receiver

	isClosed int32
}

func NewRabbitMQ(opt ConnectOption, exchangeName, exchangeType string) *RabbitMQ {
	return &RabbitMQ{
		opt:          opt,
		exchangeName: exchangeName,
		exchangeType: exchangeType,
		isClosed:     0,
	}
}

func (mq *RabbitMQ) reset() {
	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.conn = nil
	rbqlog.Printf("reset rabbitmq...")
}

func (mq *RabbitMQ) connect() error {
	var err error
	var rbqurl = fmt.Sprintf("amqp://%s:%s@%s:%d/%s", mq.opt.User, mq.opt.Password, mq.opt.Host, mq.opt.Port, mq.opt.VHost)

	mq.mu.Lock()
	defer mq.mu.Unlock()

	mq.conn, err = amqp.Dial(rbqurl)
	if err != nil {
		return err
	}

	return nil
}

func (mq *RabbitMQ) prepareExchange() error {
	if len(mq.exchangeName) == 0 {
		return nil
	}

	mq.mu.Lock()
	channel, err := mq.conn.Channel()
	mq.mu.Unlock()

	if err != nil {
		return err
	}

	err = channel.ExchangeDeclare(mq.exchangeName, mq.exchangeType, true, false, false, false,
		nil)
	if err != nil {
		return err
	}
	return nil
}

func (mq *RabbitMQ) run() {
	if err := mq.connect(); err != nil {
		rbqlog.Printf("RabbitMQ connect error.err:%s", err.Error())
		return
	}

	if err := mq.prepareExchange(); err != nil {
		rbqlog.Printf("RabbitMQ prepareExchange fail. err:%s", err.Error())
		return
	}

	for _, receiver := range mq.receivers {
		mq.wg.Add(1)
		go mq.listen(receiver)
	}

	rbqlog.Printf("RabbitMQ connect success. receivers size=%d", len(mq.receivers))

	mq.wg.Wait()
	rbqlog.Printf("RabbitMQ error, all queue has exit.")
}

func (mq *RabbitMQ) listen(receiver Receiver) {
	defer mq.wg.Done()

	queueName := receiver.QueueConfig().QueueName
	routeKey := receiver.QueueConfig().RouteKey

	mq.mu.Lock()
	channel, err := mq.conn.Channel()
	mq.mu.Unlock()

	if err != nil {
		rbqlog.Printf("RabbitMQ open channel error. %s", err.Error())
		return
	}

	_, err = channel.QueueDeclare(queueName, receiver.QueueConfig().Durable, false, receiver.QueueConfig().Exclusive, false, nil)
	if err != nil {
		rbqlog.Printf("RabbitMQ queueDeclare error.%s", err.Error())
		return
	}

	if len(mq.exchangeName) > 0 {
		err = channel.QueueBind(queueName, routeKey, mq.exchangeName, false, nil)
		if err != nil {
			rbqlog.Printf("RabbitMQ queue bind error.queueName=%s err=%s", queueName, err.Error())
			return
		}
	}

	err = channel.Qos(100, 0, true)
	if err != nil {
		rbqlog.Printf("RabbitMQ Qos error. queueName=%s err=%s", queueName, err.Error())
		return
	}

	msgChannel, err := channel.Consume(queueName, "", false, receiver.QueueConfig().Exclusive, false, false, nil)
	if err != nil {
		rbqlog.Printf("RabbitMQ Consume error. queueName=%s err=%s", queueName, err.Error())
		return
	}

	for msg := range msgChannel {
		// 重试处理，最多三次.
		retryTime := 0
		for !receiver.OnReceiver(msg) {
			retryTime++
			if retryTime > 1 {
				receiver.OnError(ErrRetry)
				break
			}
			time.Sleep(1 * time.Second)
		}
		_ = msg.Ack(false)
	}
}

func (mq *RabbitMQ) RegisterReceiver(receiver Receiver) {
	mq.receivers = append(mq.receivers, receiver)
}

func (mq *RabbitMQ) Start() {
	for atomic.LoadInt32(&mq.isClosed) == 0 {
		mq.run()
		rbqlog.Printf("RabbitMQ reconnect...")
		mq.reset()
		time.Sleep(1 * time.Second)
	}
	rbqlog.Printf("RabbitMQ exit graceful.")
}

func (mq *RabbitMQ) Close() {
	atomic.StoreInt32(&mq.isClosed, 1)

	mq.mu.Lock()
	_ = mq.conn.Close()
	mq.mu.Unlock()
}
