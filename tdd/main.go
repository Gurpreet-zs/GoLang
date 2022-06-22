package main

import (
	"strings"
)

func bob(input string) string {
	if input == "\n" {
		return "Fine. Be that way!"
	}
	if true == strings.Contains(input, "?") {
		if input == strings.ToUpper(input) {
			//	fmt.Println("1")
			return "Calm down, I know what I'm doing!"
		}
		//fmt.Println("2")
		return "Sure"
	}
	input = strings.Replace(input, " ", "", -1)
	if input == "" {
		//fmt.Println("3")
		return "Fine. Be that way!"
	}
	if input == strings.ToUpper(input) {
		//fmt.Println("4")
		return "Whoa, chill out!"
	}
	//fmt.Println("5")
	return "Whatever"
}

func main() {

}
