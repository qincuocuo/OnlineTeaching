package mkafka

import (
	"testing"
)

func Test_Consumer(t *testing.T) {
	consumer, err := NewConsumer(
		[]string{"192.168.120.20:9092"})
	if err != nil {
		t.Error(err)
		return
	}
	err = consumer.Add("assets", []int{0, 1, 2, 3, 4, 5, 6, 7}, -2)
	if err != nil {
		t.Error(err)
		return
	}
	msg := <-consumer.Output()
	t.Log(string(msg.Payload))
	consumer.Fini()
}
