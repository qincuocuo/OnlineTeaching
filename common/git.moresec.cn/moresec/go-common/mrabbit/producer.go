package mrabbit

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"

	"github.com/streadway/amqp"
)

type RBQProducer struct {
	opt ConnectOption

	mu      sync.Mutex
	conn    *amqp.Connection
	channel *amqp.Channel

	queues []QueueConfig

	msgChannel RabbitSender // 实现方保证线程安全.

	isClosed int32 // 异常原因连接不可用.
	closeCh  chan bool

	done chan bool // 成功关闭.
	wg   sync.WaitGroup
}

func NewRBQProducer(opt ConnectOption, msgChannel RabbitSender) *RBQProducer {
	if msgChannel == nil {
		msgChannel = NewDefaultChannel()
	}
	return &RBQProducer{
		opt:        opt,
		closeCh:    make(chan bool, 1),
		isClosed:   0,
		msgChannel: msgChannel,
		done:       make(chan bool, 1),
	}
}

func (rp *RBQProducer) RegisterChannel(queueList []QueueConfig) {
	rp.queues = append(rp.queues, queueList...)
}

func (rp *RBQProducer) Send(exchangeName, routeKey string, msg *amqp.Publishing) error {
	return rp.msgChannel.Write(&RabbitMsg{
		exchangeName: exchangeName,
		routeKey:     routeKey,
		body:         msg,
	})
}

func (rp *RBQProducer) prepare() error {
	var err error
	for _, q := range rp.queues {
		if len(q.ExchangeName) > 0 {
			err = rp.channel.ExchangeDeclare(q.ExchangeName, q.ExchangeType, q.Durable, false, false, false, nil)
			if err != nil {
				return err
			}
		}
		_, err = rp.channel.QueueDeclare(q.QueueName, q.Durable, false, q.Exclusive, false, nil)
		if err != nil {
			return err
		}

		if len(q.ExchangeName) > 0 {
			err = rp.channel.QueueBind(q.QueueName, q.RouteKey, q.ExchangeName, false, nil)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (rp *RBQProducer) publish(exchangeName, routeKey string, msg *amqp.Publishing) error {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	err := rp.channel.Publish(exchangeName, routeKey, false, false, *msg)
	if err != nil {
		return err
	}
	return nil
}

func (rp *RBQProducer) run() {
	err := rp.connect()
	if err != nil {
		rbqlog.Printf("RBQProducer connect error.msg:%s", err.Error())
		return
	}

	err = rp.prepare()
	if err != nil {
		rbqlog.Printf("RBQProducer declare exchange or queue error.msg:%s", err.Error())
		return
	}

	rp.wg.Add(1)
	go func() {
		defer rp.wg.Done()
	LOOP:
		for {
			select {
			case msg := <-rp.msgChannel.Read():
				err := rp.publish(msg.exchangeName, msg.routeKey, msg.body)
				if err != nil {
					rbqlog.Printf("RBQProducer send message error. routeKey:%s err:%s", msg.routeKey, err.Error())
					return
				}
			case <-rp.closeCh:
				rbqlog.Printf("RBQProducer run exit graceful.")
				break LOOP
			}
		}
		// 队列中的数据要发送完成才能推出.
		for msg := range rp.msgChannel.Read() {
			err := rp.publish(msg.exchangeName, msg.routeKey, msg.body)
			if err != nil {
				rbqlog.Printf("RBQProducer send message error. routeKey:%s err:%s", msg.routeKey, err.Error())
			}
		}
	}()

	rbqlog.Printf("RBQProducer connect success.")
	rp.wg.Wait()
}

func (rp *RBQProducer) connect() error {
	rp.mu.Lock()
	defer rp.mu.Unlock()

	var err error
	var rbqurl = fmt.Sprintf("amqp://%s:%s@%s:%d/%s", rp.opt.User, rp.opt.Password, rp.opt.Host, rp.opt.Port, rp.opt.VHost)

	rp.conn, err = amqp.Dial(rbqurl)
	if err != nil {
		return err
	}

	rp.channel, err = rp.conn.Channel()
	if err != nil {
		return err
	}
	return nil
}

func (rp *RBQProducer) Start() {
	for atomic.LoadInt32(&rp.isClosed) == 0 {
		rp.run()
		rbqlog.Printf("RBQProducer reconnect...")
		time.Sleep(2 * time.Second)
	}
}

func (rp *RBQProducer) Close() {
	close(rp.closeCh)
	atomic.StoreInt32(&rp.isClosed, 1)
}
