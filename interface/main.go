package main

import "math"

type shape interface {
	area() float64
	perimeter() float64
}
type rectangle struct {
	width  float64
	height float64
}
type circle struct {
	Radius float64
}
type square struct {
	side float64
}

type triangle struct {
	base   float64
	height float64
	side1  float64
	side2  float64
	side3  float64
}

func (r rectangle) area() float64 {
	if r.width <= 0 || r.height <= 0 {
		return -1
	}
	return r.width * r.height
}
func (c circle) area() float64 {
	if c.Radius <= 0 {
		return -1
	}
	return math.Pi * c.Radius * c.Radius
}
func (s square) area() float64 {
	if s.side <= 0 {
		return -1
	}
	return s.side * s.side
}
func (t triangle) area() float64 {
	if t.base <= 0 || t.height <= 0 || t.side1 <= 0 || t.side2 <= 0 || t.side3 <= 0 {
		return -1
	}
	return (t.base * t.height) / 2
}

func (r rectangle) perimeter() float64 {
	if r.width <= 0 || r.height <= 0 {
		return -1
	}
	return 2 * (r.width * r.height)
}
func (c circle) perimeter() float64 {
	if c.Radius <= 0 {
		return -1
	}
	return 2 * 3.14 * c.Radius
}
func (s square) perimeter() float64 {
	if s.side <= 0 {
		return -1
	}
	return 4 * s.side
}
func (t triangle) perimeter() float64 {
	if t.base <= 0 || t.height <= 0 || t.side1 <= 0 || t.side2 <= 0 || t.side3 <= 0 {
		return -1
	}
	return t.side1 + t.side2 + t.side3
}

func main() {

}
