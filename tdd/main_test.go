package main

import "testing"

func TestBob(t *testing.T) {
	testCases := []struct {
		input  string
		expOut string
	}{
		{"HOW ARE YOU", "Whoa, chill out!"},
		{"HOW ARE YOU?", "Calm down, I know what I'm doing!"},
		{"\n", "Fine. Be that way!"},
		{"How are you?", "Sure"},
		{"Hi", "Whatever"},
		{" ", "Fine. Be that way!"},
		{"", "Fine. Be that way!"},
	}

	for _, tc := range testCases {
		
		res := bob(tc.input)
		if res != tc.expOut {
			t.Error("Test Failed")
		}
	}
}
