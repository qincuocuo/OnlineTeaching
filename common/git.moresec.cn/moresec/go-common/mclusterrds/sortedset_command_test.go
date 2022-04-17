package mclusterrds

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestRedisCli_ZAdd(t *testing.T) {
	Convey("TestRedisCli_ZAdd", t, func() {
		testCli.Del("TestRedisCli_ZAdd")

		testCli.ZAdd("TestRedisCli_ZAdd", 1, "key1")
		testCli.ZAdd("TestRedisCli_ZAdd", 3, "key2")

		score, _ := testCli.ZScore("TestRedisCli_ZAdd", "key1")
		So(score, ShouldEqual, 1)

		testCli.ZIncrBy("TestRedisCli_ZAdd", 1, "key1")
		score, _ = testCli.ZScore("TestRedisCli_ZAdd", "key1")
		So(score, ShouldEqual, 2)

		rst, _ := testCli.ZRangeWithScores("TestRedisCli_ZAdd", 0, -1)
		So(rst[0].Member, ShouldEqual, "key1")
		So(rst[0].Score, ShouldEqual, 2)

		r, _ := testCli.ZRank("TestRedisCli_ZAdd", "key1")
		So(r, ShouldEqual, 0)

		r, err := testCli.ZRank("TestRedisCli_ZAdd", "keyx")
		So(r, ShouldEqual, -1)
		So(err, ShouldBeError)
	})
}
