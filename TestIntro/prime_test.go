package TestIntro

import (
	_ "math"
	"testing"
)

type TestCase struct {
	desc   string
	input  int
	expOut bool
}

func TestPrime(t *testing.T) {
	cases := []TestCase{
		{"Not-Prime", 1, false},
		{"Prime", 2, true},
		{"Prime", 11, true},
	}
	for i, tc := range cases {
		res := Prime(tc.input)
		if res != tc.expOut {
			t.Error("Number", i+1, ",Test Failed")
		}
	}
}
