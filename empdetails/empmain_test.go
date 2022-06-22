package main

import (
	"errors"
	"reflect"
	"testing"
)

func TestPost(t *testing.T) {
	testcases := []struct {
		desc   string
		input  employee
		expOut employee
		expErr error
	}{
		{"Success Case", employee{1, "Gurpreet", "Intern", 0, 0}, employee{1, "Gurpreet", "Intern", 0, 0}, nil},
		{"Invalid id Case", employee{0, "Shubham", "Intern", 0, 0}, employee{}, errors.New("Invalid Id")},
		{"Already Existing Emp", employee{1, "Jay", "Intern", 0, 0}, employee{}, errors.New("Emp Already Exists")},
		{"details of emp3", employee{3, "hello", "SE", 7, 200000}, employee{3, "hello", "SE", 7, 200000}, nil},
		{"invalid ID", employee{-2, "Joe", "HR", 10, 400000}, employee{}, errors.New("invalid ID")},
		{"ID already exists", employee{3, "Jack", "SE", 7, 200000}, employee{}, errors.New("ID already exists")},
	}
	for i, tc := range testcases {
		outPut, err := post(tc.input)
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
		expOut employee
		expErr error
	}{
		{"Success Case", 1, employee{1, "Gurpreet", "Intern", 0, 0}, nil},
		{"Emp does not exist case", 2, employee{}, errors.New("Emp does not exist")},
		{"details of emp2", 2, employee{2, "Mike", "BA", 10, 300000}, nil},
		{"details of emp3", 3, employee{3, "Jack", "SE", 7, 200000}, nil},
		{"invalid ID", -2, employee{}, errors.New("invalid ID")},
	}
	for i, tc := range testcases {
		outPut, err := get(tc.input)
		if !reflect.DeepEqual(err, tc.expErr) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
		if !reflect.DeepEqual(outPut, tc.expOut) {
			t.Errorf("%v test failed %v", i, tc.desc)
		}
	}
}
