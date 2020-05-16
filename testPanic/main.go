package main

import "fmt"

func main() {
	receivePanic()
}

func receivePanic()  {
	defer coverPanic()
	panic("i am  panic")
}

func coverPanic()  {
	message:=recover()
	fmt.Println("panic message: ", message)
}