package main

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsLeapYear(t *testing.T) {
	Convey("Positive number is given", t, func() {
		Convey("When the given number is leap year", func() {
			year := 2000

			Convey("The return value should be true", func() {
				So(IsLeapYear(year), ShouldBeTrue)
			})
		})

		Convey("When the given number is not leap year", func() {
			year := 1900

			Convey("The return value should be false", func() {
				So(IsLeapYear(year), ShouldBeFalse)
			})
		})
	})

	Convey("Negative number is given", t, func() {
		year := -100
		Convey("The return value should be false", func() {
			So(IsLeapYear(year), ShouldBeFalse)
		})
	})
}
