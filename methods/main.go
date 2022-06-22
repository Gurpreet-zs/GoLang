package main

import (
	"fmt"
	//"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Ab() float64 {
	fmt.Println("before", v.X)
	fmt.Println("before", v.Y)
	v.Y = 3
	v.X = 4
	fmt.Println("after", v.X)
	fmt.Println("after", v.Y)
	return 0.0
}

func Abs(v *Vertex) float64 {
	fmt.Println("before", v.X)
	fmt.Println("before", v.Y)
	v.Y = 7
	v.X = 8
	fmt.Println("after", v.X)
	fmt.Println("after", v.Y)
	return 0.0
}

func main() {
	v := Vertex{1, 2}
	m := Vertex{5, 6}
	fmt.Println(v.Ab())
	fmt.Println(v)
	fmt.Println(Abs(&m))
	fmt.Println(m)

}
