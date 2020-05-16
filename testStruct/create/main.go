package main

import (
	"fmt"
	"reflect"
)

type Dog struct {
	ID int
	Name string
	Age int
}


func main() {
	//testStruct1()
	//testStruct2()
	//testStruct3()
	testStruct4()
}

func testStruct1()  {
	var dog Dog
	dog.ID = 1
	dog.Name = "旺财"
	dog.Age = 5
	fmt.Println(dog)
}

func testStruct2()  {
	dog:=Dog{1,"旺财", 2}
	fmt.Println(dog)
}

func testStruct3()  {
	dog:=Dog{ID:1, Name:"旺财", Age:4}
	fmt.Println(dog)
}

func testStruct4()  {
	// 指针， 取地址操作为  & , 取对应的值为 *
	dog := new(Dog)
	dog.ID = 1
	dog.Age=13
	dog.Name="小四"
	fmt.Println(reflect.TypeOf(dog))
	fmt.Println(dog)   // &{1 小四 13}
	fmt.Println(&dog) // 0xc00000c030
	fmt.Println(*dog)  //{1 小四 13}
}
