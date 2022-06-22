package calculator

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	testCases := []struct {
		operator string
		a, b     int
		output   int
	}{
		{"+", 2, 2, 4},
		{"-", 2, 2, 0},
		{"*", 4, 4, 16},
		{"*", 10, 0, 0},
		{"/", 10, 5, 2},
		{"/", 10, 0, -1},
		{"", 12, 11, -1},
		{"   ", 12, 11, -1},
	}
	for i, tc := range testCases {
		OutCalc := calculator(tc.operator, tc.a, tc.b)
		if OutCalc != tc.output {
			t.Error("failed", i)
		}
	}

}
