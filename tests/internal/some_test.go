package internal__test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type SomeType struct {
	Key   string
	Value string
}

func TestSomething(t *testing.T) {
	Convey("For a slice", t, func() {

		s := []SomeType{
			{
				Key:   "test-key",
				Value: "test-value",
			},
			{
				Key:   "test-key",
				Value: "test-value",
			},
			{
				Key:   "test-key",
				Value: "test-value",
			},
			{
				Key:   "test-key",
				Value: "test-value",
			},
			{
				Key:   "test-key",
				Value: "test-value",
			},
			{
				Key:   "test-key",
				Value: "test-value",
			},
		}

		Convey("it should match size", func() {
			So(len(s), ShouldEqual, 6)
		})

	})
}
