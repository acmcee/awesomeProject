package main

import "fmt"

/*
... 表示可变参数
:= 声明，赋值，类型推断的功能
*/

func Sugar1(values ...string) { // 表示可变参数的，个数不一定
	for _, v := range values {
		fmt.Println("values:", v)
	}
}

func Sugar2() {
	//value := "A"
	value := false
	fmt.Println("value A:", value)
}

func main() {
	Sugar1("a", "b", "c")
	Sugar2()
}
