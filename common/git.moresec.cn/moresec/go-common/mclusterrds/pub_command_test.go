package mclusterrds

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_Subscribe(t *testing.T) {
	Convey("test pub and sub", t, func() {
		initEnv()
		result, _ := testCli.SubscribeWithTimeout("test_channel", 1*time.Second)
		fmt.Println(result)
	})
}
