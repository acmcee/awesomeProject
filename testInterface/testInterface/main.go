package main

import (
	"fmt"
)

func PrintAll(vals ...interface{}) {
	for _, v := range vals {
		if name, ok := v.(string); ok {
			fmt.Println("我是string,", name)
		} else if name, ok := v.(int); ok {
			fmt.Println("我是int,", name)
		} else {
			fmt.Printf("我是个其他类型:%T, 值是: ", v, )
			fmt.Println(v)
		}
	}

}

func TestPrint(v ...interface{}) {
	fmt.Println(v)
}

func main() {
	myArray1 := []interface{}{"1", "1112", "张三", 4, true, false, 100000000000}
	myArray2 := []interface{}{"1", "2", "张三"}
	myArray1 = append(myArray1, myArray2...)
	TestPrint(myArray1...)
}
