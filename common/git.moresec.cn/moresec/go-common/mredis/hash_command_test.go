package mredis

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_HSet(t *testing.T) {
	Convey("HSet", t, func() {
		InitEnv()
		testCli.HSet("TestRedisCli_HSet", "test", "testvalue")

		testCli.HDel("TestRedisCli_HSet", "test")
		rst, _ := testCli.HGet("TestRedisCli_HSet", "test")
		So(rst, ShouldEqual, "")

		param := map[string]interface{}{
			"test1": 1000,
			"test2": "value2",
		}
		_, _ = testCli.HMSet("TestRedisCli_HSet", param)

		mapRst, _ := testCli.HGetAll("TestRedisCli_HSet")
		So(2, ShouldEqual, len(mapRst))
		So("1000", ShouldEqual, mapRst["test1"])
		So("value2", ShouldEqual, mapRst["test2"])

		bret, _ := testCli.HExists("TestRedisCli_HSet", "test1")
		So(bret, ShouldEqual, true)
		testCli.HDel("TestRedisCli_HSet", "test1")
		bret, _ = testCli.HExists("TestRedisCli_HSet", "test1")
		So(bret, ShouldEqual, false)
		//testCli.Del("TestRedisCli_HSet")
	})
}

func TestRedisCli_HIncrBy(t *testing.T) {
	Convey("TestRedisCli_HIncrBy", t, func() {
		testCli.Del("TestRedisCli_HIncrBy")
		testCli.HIncrBy("TestRedisCli_HIncrBy", "test", 1)
		rst, _ := testCli.HGet("TestRedisCli_HIncrBy", "test")

		So(rst, ShouldEqual, "1")
		testCli.HIncrBy("TestRedisCli_HIncrBy", "test", 1)
		rst, err := testCli.HGet("TestRedisCli_HIncrBy", "test")
		So(rst, ShouldEqual, "2")

		testCli.HIncrByFloat("TestRedisCli_HIncrBy", "test2", 1.2)
		rst, err = testCli.HGet("TestRedisCli_HIncrBy", "test2")
		So(rst, ShouldEqual, "1.2")

		slist, _ := testCli.HKeys("TestRedisCli_HIncrBy")

		So(slist[0], ShouldEqual, "test")
		So(slist[1], ShouldEqual, "test2")

		l, _ := testCli.HLen("TestRedisCli_HIncrBy")
		So(l, ShouldEqual, 2)

		slist, _ = testCli.HVals("TestRedisCli_HIncrBy")
		So(slist[0], ShouldEqual, "2")
		So(slist[1], ShouldEqual, "1.2")

		mapResult, _ := testCli.HMGet("TestRedisCli_HIncrBy", []string{"test", "test2"})
		So(mapResult["test"], ShouldEqual, "2")
		So(mapResult["test2"], ShouldEqual, "1.2")

		testCli.HSetNX("TestRedisCli_HIncrBy", "test3", "value3")
		testCli.Del("TestRedisCli_HIncrBy")

		value := map[string]interface{}{
			"user_relate": 1,
			"user_id":     1323,
			"machine":     "12312434534234234jkj",
			"test1":       "12312424534234234",
			"test2":       "jdksjfksdlfjwieurisdjfslkdjfslkdfjsdfsdfwerwe",
			"test3":       "jjxkdksieusdfhjaosdlfkjfwiefjkjfskdfwiejfdlsdjfsl",
			"test4":       "jjxkdksieusdfhjaosdlfkjfwiefjkjfskdfwiejfdlsdjfsl",
		}
		testCli.HMSet("TestHashBenchmark", value)
		rsp, err := testCli.HGet("TestHashBenchmark", "user_ixd")
		t.Log(rsp, " ", err)

		rsp, err = testCli.HGet("maleware:hashfile:11", "11ac04643f2b065b24f331f86a804aeb")
		t.Log(rsp, " ", err)
	})
}
