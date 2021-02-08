package main

import (
	"fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestImpl(t *testing.T) {
	tests := getTests()

	for _, test := range tests {
		desc := fmt.Sprintf("Computing the infix expression: %s", test.input)
		Convey(desc, t, func() {
			result := compute(test.input)
			Convey(fmt.Sprintf("Should equal: %f", test.want), func() {
				So(result, ShouldEqual, test.want)
			})
		})
	}
}

type test struct {
	input string
	want  float64
}

func getTests() []test {
	return []test{
		{
			input: "(20-(8-3))-3",
			want:  12,
		},
		{
			input: "1",
			want:  1,
		},
		{
			input: "(1)",
			want:  1,
		},
		{
			input: "(1+2)",
			want:  3,
		},
		{
			input: "1+2",
			want:  3,
		},
		{
			input: "(1+2)*4/2^2^3",
			want:  0.1875,
		},
	}
}
