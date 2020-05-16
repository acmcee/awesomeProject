package goroutine

import (
	"fmt"
	"time"
)

func TestLoop()  {
	for i:=0;i<11 ;i++  {

		fmt.Printf("%d, ", i)
		time.Sleep(time.Microsecond * 100)
	}
	fmt.Println()
}