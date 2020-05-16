package main

import "fmt"

type Dog struct {
	ID int
	Name string
	Age int
}

func (d *Dog)Run()  {  // 如果要被其他的包使用的话，必须是大写的哦，否则不能被其他包引用
	fmt.Println("我是", d.Name, ",Dog 正在跑步哦")
}

func CreateDog() *Dog {

	d := new(Dog)
	d.Name = "旺财"
	d.ID = 1
	d.Age = 5
	return d
}

func main() {
	d:= CreateDog()
	d.Run()
}
