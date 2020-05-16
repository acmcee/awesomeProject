package main

import (
	"awesomeProject/testGoroutine/goroutine"
	"time"
)

func main() {
	/*fmt.Println("cpu num: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	time.Sleep(time.Second)*/


	go goroutine.TestSend()
	go goroutine.TestReceive()
	time.Sleep(time.Second* 10)

}
