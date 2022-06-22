package main

import (
	"reflect"
	"testing"
)

func TestEmployee(t *testing.T) {
	testCases := []struct {
		input  employeeDetails
		expOut employeeDetails
	}{
		{employeeDetails{"Jay", "19/01/2000", 0, 1, 0}, employeeDetails{"Jay", "19/01/2000", 22, 1, 20000}},
		{employeeDetails{"Shubham", "12/02/2000", 0, -1, 0}, employeeDetails{"Shubham", "12/02/2000", 22, 2, 50000}},
		{employeeDetails{"Raj", "22/01/2000", 0, 4, 0}, employeeDetails{"Raj", "22/01/2000", 22, 4, 50000}},
		{employeeDetails{"Rohan", "", 0, 6, 0}, employeeDetails{"Rohan", "", 22, 6, 100000}},
		{employeeDetails{"Sham", "19/01/1990", 0, 10, 0}, employeeDetails{"Sham", "19/01/2000", 32, 1, 200000}},
		{employeeDetails{" ", "23/12/2003", 0, 3, 0}, employeeDetails{"rj", "23/12/2003", 19, 3, 50000}},
		{employeeDetails{"", "23/12/2003", 0, 3, 0}, employeeDetails{"rj", "23/12/2003", 19, 3, 50000}},
	}
	for _, tc := range testCases {
		if !reflect.DeepEqual(tc.input, tc.expOut) {
			t.Error("test Failed")
		}
	}
}
