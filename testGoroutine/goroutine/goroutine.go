package goroutine

import (
	"fmt"
	"time"
)


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