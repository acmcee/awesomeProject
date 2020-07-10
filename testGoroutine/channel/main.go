package main

import (
	"fmt"
	"time"
)

func ArraySource() <-chan int {

	ch := make(chan int, 100) // 由于这里是创建了无缓冲的channel，因此必须有个用goroutine 去写数据，否则就是deadlock了
	// 否则就是写的时候堵塞，后面的读将无法执行到，main 卡死
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch2 := ArraySource()

	for {
		if d, ok := <-ch2; ok { // 也可以用ok 来判断，chan 是否被关闭掉
			fmt.Println("ch2", d)
			fmt.Println("len", len(ch2))
			fmt.Println("cap", cap(ch2))
		} else {
			fmt.Println("ch closed")
			break
		}
		time.Sleep(time.Second)

	}

}
