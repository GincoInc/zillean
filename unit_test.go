package zillean

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestUnit_FromQa(t *testing.T) {
	Convey("converts a number in Qa to that in specified unit", t, func() {
		result, err := FromQa("1000000000000", "zil", false)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "1")
		result, err = FromQa("1000000", "li", false)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "1")
		result, err = FromQa("-1000000000000", "zil", false)
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "-1")
	})
}

func TestUnit_ToQa(t *testing.T) {
	Convey("converts a number in Qa to that in specified unit", t, func() {
		result, err := ToQa("1", "li")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "1000000")
		result, err = ToQa("1", "zil")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "1000000000000")
		result, err = ToQa("-1", "zil")
		So(err, ShouldBeNil)
		So(result, ShouldEqual, "-1000000000000")
	})
}
