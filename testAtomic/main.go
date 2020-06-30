package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {

	var a int32 = 10

	var wg sync.WaitGroup
	bTime := time.Now()
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt32(&a, 10)

			wg.Done()

		}()
	}
	wg.Wait()
	fmt.Printf("expect %d, really %d , cost %v", 10000*10+10, atomic.LoadInt32(&a), time.Since(bTime))
}
