package main

import (
	"fmt"
)

func main() {
	receivePanic()
}

func receivePanic()  {

	defer coverPanic()
	//panic("i am  panic")
	//panic(errors.New("我是个error 类型"))
	panic(1)
}

func coverPanic()  {
	message:=recover()
	switch message.(type) {  // 根据类型去捕获异常
	case string:
		fmt.Println("string message: ", message)
	case error:
		fmt.Println("error message: ", message)
	default:
		fmt.Println("unknow message:", message)
	}
}