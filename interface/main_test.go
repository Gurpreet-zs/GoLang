package main

import (
	"testing"
)

func TestArea(t *testing.T) {

	areaTests := []struct {
		desc   string
		shape  shape
		output float64
	}{
		{"rectangle test", rectangle{12, 6}, 72.0},
		{"rectangle With -ve values", rectangle{-12, 0}, -1},
		{"square test", square{4}, 16.0},
		{"square with -ve values", square{-29}, -1},
		{"circle test", circle{10}, 314.1592653589793},
		{"circle with -ve values", circle{-33}, -1},
		{"triangle test", triangle{12, 6, 2, 2, 2}, 36.0},
		{"triangle With -ve test(base)1", triangle{-12, 6, 2, 2, 2}, -1},
		{"triangle With -ve test(side)", triangle{12, 6, 2, -2, 2}, -1},
		{"triangle With noSide test(base)", triangle{12, 6, 0, 2, 2}, -1},
	}

	for i, tc := range areaTests {
		op := tc.shape.area()
		if op != tc.output {
			t.Errorf("testCase %d failed: %s", i, tc.desc)
		}
	}

}

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		desc   string
		shape  shape
		output float64
	}{
		{"rectangle test", rectangle{12, 6}, 36},
		{"rectangle With -ve values", rectangle{-12, 0}, -1},
		{"square test", square{4}, 16.0},
		{"square with -ve values", square{-29}, -1},
		{"circle test", circle{10}, 62.83},
		{"circle with -ve values", circle{-33}, -1},
		{"triangle test", triangle{12, 6, 12, 2, 2}, 16},
		{"triangle test with base 0", triangle{0, 6, 12, 2, 2}, -1},
		{"triangle With -ve test(base)", triangle{-12, 6, 2, 2, 2}, -1},
		{"triangle With -ve test(side)", triangle{12, 6, 2, -2, 2}, -1},
		{"triangle With noSide test(base)", triangle{12, 6, 0, 2, 2}, -1},
	}

	for i, tc := range perimeterTests {
		op := tc.shape.perimeter()
		if op != tc.output {
			t.Errorf("testCase %d failed: %s", i, tc.desc)
		}
	}

}
