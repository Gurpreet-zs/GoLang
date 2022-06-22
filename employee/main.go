package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type employeeDetails struct {
	name   string
	dob    string
	age    int
	exp    int
	salary int
}

func (emp employeeDetails) greeting() {
	log.Println("Invalid Experience")
	res := strconv.Itoa(emp.exp)
	fmt.Println("Hello " + emp.name + " Congratulations for completing " + res + " Years")

}

func (emp *employeeDetails) getAge() string {
	if emp.dob == "" {
		return "Invalid Date"
	}
	res := strings.Split(emp.dob, "/")
	//res1, err := strconv.ParseInt(res[2], 2, 64)
	int1, err := strconv.ParseInt(res[2], 10, 64)
	if err != nil {
		log.Println(err)
		return "Error in date"
	}
	years := 2022 - int1
	emp.age = int(years)

	return " "
}

func (emp *employeeDetails) getSalary() string {
	if emp.exp < 0 {
		log.Println("Invalid Experince")
	} else if emp.exp > 0 || emp.exp <= 1 {
		emp.salary = 20000
		//fmt.Println(emp.greeting())
	} else if emp.exp >= 2 || emp.exp <= 5 {
		emp.salary = 50000
		//fmt.Println(emp.greeting())
	} else if emp.exp > 5 || emp.exp <= 10 {
		emp.salary = 100000
		//fmt.Println(emp.greeting())
	} else if emp.salary > 10 {
		emp.salary = 200000
		//fmt.Println(emp.greeting())
	} else {
		emp.salary = 0
		emp.age = 0
		emp.dob = ""
		emp.name = ""
		emp.exp = 0
		log.Println("Invalid Experience")
		return ""

	}

	return ""
}
func empResult(e employeeDetails) employeeDetails {
	e.getSalary()
	e.getAge()
	return e

}
func main() {

}
