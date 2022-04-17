package mclusterrds

import (
	"fmt"
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	initEnv()
}

func TestRedisCli_Append(t *testing.T) {
	Convey("test append", t, func() {
		testCli.Del("TestRedisCli_Append")
		defer testCli.Del("TestRedisCli_Append")

		ok, err := testCli.Set("TestRedisCli_Append", "test")
		So(ok, ShouldEqual, 1)
		So(err, ShouldBeNil)

		result, err := testCli.Get("TestRedisCli_Append")
		So(result, ShouldEqual, "test")
		So(err, ShouldBeNil)

		result, err = testCli.Get("TestRedisCli_Append_Nil")
		So(result, ShouldBeEmpty)
		So(err, ShouldBeError)
	})
}

func TestRedisCli_Decrby(t *testing.T) {
	Convey("test decrby", t, func() {
		testCli.Del("TestRedisCli_Decrby")
		defer testCli.Del("TestRedisCli_Decrby")

		for i := 1; i < 4; i++ {
			result, err := testCli.Decr("TestRedisCli_Decrby")
			So(err, ShouldBeNil)
			So(result, ShouldEqual, -1*i)
		}

		result, err := testCli.Incr("TestRedisCli_Decrby")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, -2)
		result, err = testCli.IncrBy("TestRedisCli_Decrby", 2)
		So(result, ShouldEqual, 0)

		result, err = testCli.IncrBy("TestRedisCli_Decrby", 2100000000)
		So(result, ShouldEqual, 2100000000)
	})
}

func TestRedisCli_Exists(t *testing.T) {
	Convey("test_exists", t, func() {
		testCli.Del("TestRedisCli_Exists")
		defer testCli.Del("TestRedisCli_Exists")

		testCli.Set("TestRedisCli_Exists", "test")
		rt, err := testCli.Exists("TestRedisCli_Exists")
		So(rt, ShouldBeTrue)
		So(err, ShouldBeNil)

		testCli.Del("TestRedisCli_Exists")
		rt, err = testCli.Exists("TestRedisCli_Exists")
		So(rt, ShouldBeFalse)
		So(err, ShouldBeNil)
	})
}

func TestRedisCli_TTL(t *testing.T) {
	Convey("test_ttl", t, func() {
		testCli.Del("TestRedisCli_TTL")
		defer testCli.Del("TestRedisCli_TTL")

		testCli.SetEx("TestRedisCli_TTL", "test", 100)
		rt, err := testCli.TTL("TestRedisCli_TTL")
		fmt.Println("rt=", rt)
		So(rt, ShouldEqual, 100)
		So(err, ShouldBeNil)

		testCli.Set("TestRedisCli_TTL_no", "test")
		rt, err = testCli.TTL("TestRedisCli_TTL_no")
		So(rt, ShouldEqual, -1)
		So(err, ShouldBeNil)

		rt, err = testCli.TTL("TestRedisCli_TTL_err")
		So(rt, ShouldEqual, -2)
		So(err, ShouldBeNil)

		testCli.Del("TestRedisCli_NX_EX")
		rtb, err := testCli.SetNX("TestRedisCli_NX_EX", "value", 100*time.Second)
		So(rtb, ShouldBeTrue)
		So(err, ShouldBeNil)

		rt, err = testCli.TTL("TestRedisCli_NX_EX")
		So(rt, ShouldEqual, 100)
		So(err, ShouldBeNil)
	})
}
