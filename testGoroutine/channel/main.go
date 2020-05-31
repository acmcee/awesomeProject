package main

import "fmt"

func ArraySource(a ...int) <-chan int {

	ch := make(chan int) // 由于这里是创建了无缓冲的channel，因此必须有个用goroutine 去写数据，否则就是deadlock了
	// 否则就是写的时候堵塞，后面的读将无法执行到，main 卡死
	go func() {
		for i := range a {
			ch <- i
		}
		close(ch)
	}()
	return ch
}

func main() {
	ch1 := ArraySource(1, 2, 3, 4, 5, 6)
	ch2 := ArraySource(1, 2, 3, 4, 5, 6)
	for i := range ch1 { // 如果chan close 的话，可以直接用range
		fmt.Println("ch1", i)
	}

	for {
		if d, ok := <-ch2; ok { // 也可以用ok 来判断，chan 是否被关闭掉
			fmt.Println("ch2", d)
		} else {
			break
		}
	}

}
