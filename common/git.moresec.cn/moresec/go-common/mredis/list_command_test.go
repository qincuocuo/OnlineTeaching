package mredis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_ListOption(t *testing.T) {
	InitEnv()
	Convey("test", t, func() {
		testCli.Del("TestRedisCli_ListOption")
		iLen, _ := testCli.ListLen("TestRedisCli_ListOption")
		So(iLen, ShouldEqual, 0)

		testCli.ListLPush("TestRedisCli_ListOption", "value1")
		testCli.ListLPush("TestRedisCli_ListOption", "value2")
		testCli.ListLPush("TestRedisCli_ListOption", "value3")
		sRet, _ := testCli.ListIndex("TestRedisCli_ListOption", -1)
		So(sRet, ShouldEqual, "value1")

		sRet, _ = testCli.ListLPop("TestRedisCli_ListOption")
		So(sRet, ShouldEqual, "value3")

		testCli.Del("TestRedisCli_ListOption")

		dataStr := []string{"1", "2", "3"}
		_, err := testCli.ListRPush("testRpush", dataStr)
		So(err, ShouldBeNil)
		testCli.Del("testRpush")
	})
}

func TestRedisCli_ListRPopLPush(t *testing.T) {
	InitEnv()
	Convey("test", t, func() {
		testCli.Del("TestRedisCli_ListRPopLPush")
		testCli.Del("TestRedisCli_dest")

		testCli.ListLPush("TestRedisCli_ListRPopLPush", []string{"a", "b", "c"})
		testCli.ListRPopLPush("TestRedisCli_ListRPopLPush", "TestRedisCli_ListRPopLPush")
		result, err := testCli.ListLRange("TestRedisCli_ListRPopLPush", 0, -1)
		So(err, ShouldBeNil)
		So(result[0], ShouldEqual, "a")
		So(result[1], ShouldEqual, "c")
		So(result[2], ShouldEqual, "b")

		testCli.ListRPopLPush("TestRedisCli_ListRPopLPush", "TestRedisCli_dest")
		result, err = testCli.ListLRange("TestRedisCli_dest", 0, -1)
		So(err, ShouldBeNil)
		So(result[0], ShouldEqual, "b")
	})
}
