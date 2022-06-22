package main

import "fmt"

//type I interface {
//	M()
//}
//type T struct {
//	S string
//}
//
//func (t *T) M() {
//	if t == nil {
//		fmt.Println("nil")
//		return
//	}
//	fmt.Println(t.S)
//
//}

var sum int

func add(i int) {
	sum += i
	fmt.Println(sum, i)
}

func main() {

	fmt.Println("counting")
	defer fmt.Println(sum)
	for i := 0; i < 10; i++ {
		defer add(i)
	}
	defer fmt.Println(sum)

	//i := 0
	//defer fmt.Println("start")
	//defer fmt.Println("defer 1 i : ", i)
	//
	//i++
	//defer fmt.Println("defer 2 i :", i)
	//defer fmt.Println("end")

	//var i I
	//var t *T
	//i = t
	//
	//fmt.Printf("(%v, %T)\n", i, i)
	//i.M()
	//
	//i = &T{"hello man"}
	//fmt.Printf("(%v, %T)\n", i, i)
	//i.M()
}
