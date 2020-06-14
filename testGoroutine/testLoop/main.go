package main

import (
	"fmt"
	"math/rand"
	"time"
)

func TestMakeCake(i int, doneCh chan<- int, cooktime time.Duration) {
	fmt.Println(i, "号蛋糕要做", cooktime)
	time.Sleep(cooktime)
	fmt.Println(i, "号蛋糕做好了")
	doneCh <- i
}

func main() {
	makeNum := 10
	ch := make(chan int, 2)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < makeNum; i++ {

		go TestMakeCake(i, ch, time.Second*time.Duration(rand.Int()%5))
	}

	for i := 0; i < makeNum; i++ {

		fmt.Println("我知道", <-ch, "号蛋糕做好了")
	}

	fmt.Println("全部蛋糕做好了")
}
