package main

import (
	"context"
	"fmt"
	"time"
)

func TestTick(ctx context.Context) {
	t := time.Tick(time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("TestTick cancel 啦...")
			return
		case <-t:
			fmt.Println("TestTick tick ....")
		}
	}
}

func main() {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(time.Second*5))
	defer cancel()
	go TestTick(ctx)

	time.Sleep(time.Second * 10)

}
