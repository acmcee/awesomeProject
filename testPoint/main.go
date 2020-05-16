package main

import "fmt"

/*
指针的基本使用
指针数组
指向指针的指针
值传递和指针传递



* 除了可以用声明指针类型，也是取值运算符
& 是取对应的指针的用处

 */

var name int
var namePoint *int
var nullPoint *int

/*

指针数组 ，就是个数组，里面存了指针

数组指针
 */

func TestPoint() {
	name = 1
	namePoint = &name // 取对应的变量的指针
	fmt.Println("name:", name)
	fmt.Println("&name:", &name)
	fmt.Println("namePoint:", namePoint)
	fmt.Println("&namePoint:", &namePoint)
	fmt.Println("namePoint指向的值是:", *namePoint)
	fmt.Println("nullPoint:", nullPoint)
	//fmt.Println("*nullPoint:", *nullPoint)  // 这个不能打印，指针没有赋值的话，对应的指向是nil
}

func TestPintArray() {
	a, b := 1, 2
	pointArr := [...]*int{&a, &b} // 表示是个指针数据  [2]*int 表示 2个长度的数组，里面存的是指针，指向的是int类型

	fmt.Println(pointArr) // [0xc0000140a0 0xc0000140a8]

}

func TestArrayPoint() {
	arr := [...]int{1, 2, 3}
	fmt.Println(&arr) // &[1 2 3]

}

func main() {
	// 测试指针的使用，获取指针的值，以及取对应的指针
	//TestPoint()

	// 指针数组
	TestPintArray()

	// 数组的指针
	TestArrayPoint()
}
