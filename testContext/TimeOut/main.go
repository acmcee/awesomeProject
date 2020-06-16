package main

import (
	"context"
	"fmt"
	"time"
)

func TestTick(ctx context.Context) {
	t := time.NewTicker(time.Second)
	t.Stop()
	for {

		select {
		case <-ctx.Done():
			fmt.Println("cancel 啦...")

			return
		case <-t.C:

			fmt.Println("tick ....")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	go TestTick(ctx)

	time.Sleep(time.Second * 6)
	cancel()
	time.Sleep(time.Second * 2)

}
