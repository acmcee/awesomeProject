package main

import (
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	Name      string
	Address   string
	DoB       time.Time
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func EmployeeByID(id int) *Employee {
	/* ... */
	e := new(Employee)
	e.ID = id
	e.Position = "aaaaaaaaaaaaa"
	return e
}

func main() {

	fmt.Println(EmployeeByID(dilbert.ManagerID).Position) // "Pointy-haired boss"

	id := dilbert.ID
	// 如果  EmployeeByID 返回的是 Employee 的话，下面语句将报错， 是个值类型
	// EmployeeByID(id).Position = "aaa"  //  Cannot assign to EmployeeByID(id).Position

	// 返回的是指针变量，我们可以直接操作指针的数据
	EmployeeByID(id).Position = "aaa"

}
