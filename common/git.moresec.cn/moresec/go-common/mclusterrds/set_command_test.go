package mclusterrds

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_SAdd(t *testing.T) {
	initEnv()
	Convey("TestRedisCli_SAdd", t, func() {
		testCli.Del("TestRedisCli_SAdd")

		testCli.SAdd("TestRedisCli_SAdd", []string{"value1", "value2", "value3"})

		iCount, _ := testCli.SCard("TestRedisCli_SAdd")
		So(iCount, ShouldEqual, 3)

		bRet, _ := testCli.SIsMember("TestRedisCli_SAdd", "value1")
		So(bRet, ShouldEqual, true)

		bRet, _ = testCli.SIsMember("TestRedisCli_SAdd", "valuex")
		So(bRet, ShouldEqual, false)

		So(3, ShouldEqual, 3)
		//testCli.Del("TestRedisCli_SAdd")
	})
}
