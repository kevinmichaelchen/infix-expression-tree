package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestImpl(t *testing.T) {
	type test struct {
		input string
		want  int
	}
	tests := []test{
		{
			input: "(20- (8-3)) - 3",
			want:  12,
		},
	}

	for _, test := range tests {
		desc := fmt.Sprintf("Computing the infix expression: %s", test.input)
		Convey(desc, t, func() {
			result := compute(test.input)
			Convey(fmt.Sprintf("Should equal: %d", test.want), func() {
				So(result, ShouldEqual, test.want)
			})
		})
	}
}
