package mkafka

import (
	"fmt"
	"testing"
	"time"
)

//生产端#go test -v -run=Test_Producer
//消费端kafka-console-consumer --bootstrap-server localhost:9092 --topic singchia_produce_test
func Test_Producer(t *testing.T) {
	failedCh := make(chan *ProducerMessage)
	go func() {
		for failedMsg := range failedCh {
			t.Logf("failed msg: %s, err: %s", failedMsg.Payload, failedMsg.Error)
		}
	}()

	succeedCh := make(chan *ProducerMessage)
	go func() {
		for succeedMsg := range succeedCh {
			t.Logf("succeed msg: %s", succeedMsg.Payload)
		}
	}()

	producer, err := NewProducer(
		[]string{"192.168.120.20:9092"},
		OptionFailedCh(failedCh),
		OptionSucceedCh(succeedCh))
	if err != nil {
		t.Error(err)
		return
	}
	template1 := `
{
	"msg_type": "host",
	"ip_addr": "192.168.199.56",
	"ts": %d,
}
	`

	template2 := `
{
	"msg_type": "host_port",
	"ip_addr": "192.168.199.56",
	"port": 80,
	"cnt": "XEFADg14=",
	"ts": %d,
}
`

	template3 := `
{
	"msg_type": "host_domain",
	"ip_addr": "192.168.199.197",
	"host": "git.moresec.cn",
	"ts": %d,
}
	`

	asset1 := fmt.Sprintf(template1, time.Now().Unix())
	asset2 := fmt.Sprintf(template2, time.Now().Unix())
	asset3 := fmt.Sprintf(template3, time.Now().Unix())
	producer.Input() <- &ProducerMessage{
		Topic:   "assets",
		Payload: []byte(asset1),
	}
	producer.Input() <- &ProducerMessage{
		Topic:   "assets",
		Payload: []byte(asset2),
	}
	producer.Input() <- &ProducerMessage{
		Topic:   "assets",
		Payload: []byte(asset3),
	}
	producer.Fini()
}
