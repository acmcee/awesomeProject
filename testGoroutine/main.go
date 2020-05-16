package main

import (
	"awesomeProject/testGoroutine/goroutine"
	"fmt"
)

func main() {
	/*fmt.Println("cpu num: ", runtime.NumCPU())
	runtime.GOMAXPROCS(runtime.NumCPU() - 1)
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	time.Sleep(time.Second)*/

	/*
		// 测试携程数据的同步，利用select
		go goroutine.TestSend()
		go goroutine.TestReceive()
		time.Sleep(time.Second* 10)
	*/

	// 测试携程的同步
	goroutine.TestRead()
	go goroutine.TestWrite()
	goroutine.WG.Wait()
	fmt.Println("ALL Done")

}
