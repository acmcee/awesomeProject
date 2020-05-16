package main

import (
	"awesomeProject/testGoroutine/goroutine"
	"time"
	"runtime"
	"fmt"
)

func main() {
	fmt.Println("cpu num: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	time.Sleep(time.Second)
}
