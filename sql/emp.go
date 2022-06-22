package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
)


type Employee struct {
	empID     int
	empName   string
	empPos    string
	empExp    int
	empSalary int
}

func GetDetails(emp Employee) int {
	return emp.empID
}
func PostDetails(emp Employee) (int, int) {
	return emp.empID, 0
}

package main

import (
"database/sql"
"fmt"
"log"

_ "github.com/go-sql-driver/mysql"
)

type employee_info struct {
	id                 int
	name, position     string
	experience, salary float64
}

func CreateTable() {
	db, err := sql.Open("mysql", "root:Gurpreet@0848-03@/")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(db.Ping())
	_, err = db.Exec("CREATE DATABASE test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created")
	}
	_, err = db.Exec("USE test")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("database selected successfully")
	}
	table, err := db.Prepare("CREATE Table employee(emp_id int NOT NULL ,emp_name varchar(50),emp_position varchar(50),emp_experience float,emp_salary float,PRIMARY KEY (emp_id));")
	if err != nil {
		fmt.Println(err.Error())
	}
	_, err = table.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("table created successfully")
	}

}

func PostDetails(emp Employee) employee_info {
	db, err := sql.Open("mysql", "root:Gurpreet@0848-03@/test")
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	sql := "INSERT INTO employee() VALUES ()"

	res, err := db.Exec(sql)

	if err != nil {
		panic(err.Error())
	}

	lastId, err := res.LastInsertId()

}

func GetData(i int) employee_info {

}