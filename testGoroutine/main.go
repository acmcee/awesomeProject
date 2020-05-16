package main

import (
	"awesomeProject/testGoroutine/goroutine"
	"time"
)

func main() {
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	go goroutine.TestLoop()
	time.Sleep(time.Second)
}
