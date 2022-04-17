package mpkg

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestCompareVersion(t *testing.T) {
	Convey("test apt version", t, func() {
		So(CompareVersion("4.3-11+deb8u1", "4.3-11"), ShouldEqual, 1)
		So(CompareVersion("4.3-11+deb8u1", "4.3-11+deb8u1"), ShouldEqual, 0)
		So(CompareVersion("4.3-11+deb8u1", "4.3-12"), ShouldEqual, -1)
		So(CompareVersion("2.11.e1", "2.12"), ShouldEqual, -1)
		So(CompareVersion("0A.3", "000B"), ShouldEqual, -1)
		So(CompareVersion("2.11", "3"), ShouldEqual, -1)
		So(CompareVersion("1.8.4-5ubuntu1.2", "1.8.4-5ubuntu1.3"), ShouldEqual, -1)
	})
}
