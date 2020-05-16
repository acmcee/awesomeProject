package main

import "fmt"

type Behavior interface {
	Run() string
	Eat() string
}

type Animal struct {
	Color string
}

type Dog struct {
	Animal
	ID int
	Name string
	Age int
}

type Cat struct {
	Animal
	ID int
	Name string
	Age int
}


func (d *Dog)Eat() string {  // 已经隐式实现了接口
	fmt.Println(d.Name, " dog is eatting")
	return "eating"
}

func (d *Dog)Run()  string { 	// 已经隐式实现了接口
	fmt.Println(d.Name, " dog is runing")

	return "running"
}




func (c *Cat)Eat() string {  // 已经隐式实现了接口
	fmt.Println(c.Name, " cat is eatting")
	return "eating"
}

func (c *Cat)Run()  string { 	// 已经隐式实现了接口
	fmt.Println(c.Name, " cat is runing")
	return "running"
}


// 接口方法
func action(d Behavior) {
	d.Run()
	d.Eat()
}



func main() {

	// 测试接口定义变量

	//var b Behavior  // 定义变量为接口
	//b = new(Dog)    // 用Dog 给他赋值

	//b.Run()

	action(new(Cat))
	action(new(Dog))

	
}
