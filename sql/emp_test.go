package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	testcases := []struct {
		desc   string
		input  Employee
		expOut Employee
		expErr error
	}{
		{"Success Case", Employee{1, "Jay", "Intern", 0, 0}, Employee{1, "Jay", "Intern", 0, 0}, nil},
		{"Invalid id Case", Employee{0, "Shubham", "Intern", 0, 0}, Employee{}, errors.New("Invalid Id")},
		{"Already Existing Emp", Employee{1, "Jay", "Intern", 0, 0}, Employee{}, errors.New("Emp Already Exists")},
	}
	for i, tc := range testcases {
		outPut, err := Post(tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		if !reflect.DeepEqual(outPut, tc.expOut) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}

func TestGet(t *testing.T) {
	testcases := []struct {
		desc   string
		input  int
		expOut Employee
		expErr error
	}{
		{"Success Case", 1, Employee{1, "Jay", "Intern", 0, 0}, nil},
		{"Emp does not exist case", 2, Employee{}, errors.New("Emp does not exist")},
	}
	for i, tc := range testcases {
		outPut, err := Get(tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		if !reflect.DeepEqual(outPut, tc.expOut) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}
