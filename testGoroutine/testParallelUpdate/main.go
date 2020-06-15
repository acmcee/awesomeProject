package main

import (
	"fmt"
	"sync"
	"time"
)

var balance1 = 100
var balance2 = 100
var mu sync.Mutex

func UpdateBalance1() {
	balance1 += 10
}

func UpdateBalance2() {
	mu.Lock()
	balance2 += 10
	mu.Unlock()
}

func main() {
	var wg sync.WaitGroup
	bTime := time.Now()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			UpdateBalance1()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("balance1 expect ", 100+1000*10, ", now is:", balance1, "cost:", time.Since(bTime))
	bTime = time.Now()

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			UpdateBalance2()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("balance2 expect ", 100+1000*10, ", now is:", balance2, "cost:", time.Since(bTime))

	//balance1 expect  10100 , now is: 9610 cost: 884.253µs
	//balance2 expect  10100 , now is: 10100 cost: 339.304µs

	// 实际发现第二个更快，第二个还加锁。。，第一个记录会更新丢失，第二个原子更新

}
