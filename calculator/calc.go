package calculator

import (
	"log"
	"strings"
)

func calculator(op string, a int, b int) int {
	op = strings.TrimSpace(op)
	if op == "" {
		log.Println("Invalid Operator")
		return -1
	}

	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			return -1
		} else {
			return a / b
		}

	}
	return -1
}
