package goroutine

import (
	"fmt"
	"time"
	"sync"
)

var WG sync.WaitGroup

var mChan = make(chan int, 5)
var timeoutChan = make(chan bool)
func TestLoop()  {
	for i:=0;i<11 ;i++  {

		fmt.Printf("%d, ", i)
		time.Sleep(time.Microsecond * 100)
	}
	fmt.Println()
}

func TestSend()  {
	for i:=0;i<5 ;i++  {
		mChan <- i
		fmt.Println("我发送了: ", i)
		time.Sleep(time.Millisecond* 10)
	}

	timeoutChan <- true
}

func TestReceive()  {
	/*for i:=0;i<5 ;i++  {
		num:= <- mChan
		fmt.Println("我收到了: ", num)
	}*/

	for  {
		select {

		case num:= <-mChan:
			fmt.Println("我收到了：", num)
		case <- timeoutChan:
			fmt.Println("我收到了超时的Flag")

		}
	}
}

func TestRead() {
	for i := 0; i < 3; i++ {
		WG.Add(1)
	}
}

func TestWrite() {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second)
		fmt.Println("GET : ", i)
		WG.Done()
	}
}

func TestMultiChan() {
	messages1 := make(chan int, 10)
	messages2 := make(chan int, 10)
	messages3 := make(chan int, 10)
	done := make(chan bool)

	defer close(messages1)
	defer close(messages2)
	defer close(messages3)
	// consumer
	go func() {
		ticker := time.NewTicker(1 * time.Millisecond * 100)
		for _ = range ticker.C {
			select {
			case <-done:
				fmt.Println("child process interrupt...")
				return
			case d := <-messages1:
				fmt.Printf("get from message1: %d\n", d)
			case d := <-messages2:
				fmt.Printf("get from message2: %d\n", d)
			case d := <-messages3:
				fmt.Printf("get from message3: %d\n", d)
			}
		}
	}()

	// producer
	for i := 0; i < 10; i++ {
		messages1 <- i + 100
		messages2 <- i + 200
		messages3 <- i + 300
	}
	time.Sleep(5 * time.Second)
	close(done)
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
}
