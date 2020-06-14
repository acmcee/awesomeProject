package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func MakeCake(i int, doneCh chan<- int) {
	rand.Seed(time.Now().UnixNano())
	cooktime := time.Second * time.Duration(rand.Int()%5)
	fmt.Println(i, "号蛋糕要做", cooktime)
	time.Sleep(cooktime)
	fmt.Println(i, "号蛋糕做好了")
	doneCh <- i
}

func TestMakeCake1() {
	makeNum := 10
	ch := make(chan int, 2)
	for i := 0; i < makeNum; i++ {
		go MakeCake(i, ch)
	}
	for i := 0; i < makeNum; i++ {
		fmt.Println("我知道", <-ch, "号蛋糕做好了")
	}
	fmt.Println("全部蛋糕做好了")
}

func TestMakeCake2() {
	makeNum := 10
	ch := make(chan int, 2)
	var wg sync.WaitGroup
	for i := 0; i < makeNum; i++ {
		wg.Add(1)
		go func(i int) {
			MakeCake(i, ch)
			wg.Done()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for i := range ch {
		fmt.Println("我知道", i, "做好了")
	}

	fmt.Println("全部蛋糕做好了")
}

func main() {
	TestMakeCake2()
}
