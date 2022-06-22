package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("Twice %q is %v\n", v, len(v))
	default:
		fmt.Printf("i dont know %T\n", v)
	}

}
func main() {
	do(21)
	do("heyy")
	do(true)
	do(23.23)

}
