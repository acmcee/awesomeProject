package main

import "fmt"

/*
组合，面向对象的特性，继承，各种结构体互相组合
 */

type Animal struct {
	Color string
}

type Dog struct {
	Animal
	ID int
	Name string
	Age int
}

func (d *Dog) Run()  {
	fmt.Println("我是一只狗，我会跑，我的名字是", d.Name)
}

func (d *Dog)Eat()  {
	fmt.Println(d.Name, "yummy....delicious")
}

func (d *Animal)LoveColor()  {
	fmt.Println("我喜欢的颜色是" , d.Color)
}

func main() {
	dog := new(Dog)
	dog.Name = "旺财"
	dog.Age = 5
	dog.ID = 1
	dog.Color = "红色"
	dog.Run()
	dog.Eat()
	dog.LoveColor()
}
