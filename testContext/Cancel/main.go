package main

import (
	"context"
	"fmt"
	"time"
)

func TestTickChild(ctx context.Context) {
	t := time.Tick(time.Second)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("TestTickChild cancel 啦...")
			return
		case <-t:
			fmt.Println("TestTickChild tick ....")
		}
	}
}

func TestTick(ctx context.Context) {
	t := time.Tick(time.Second)
	newCtx, cancel := context.WithCancel(ctx)
	defer cancel()
	go TestTickChild(newCtx)

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
	ctx, cancel := context.WithCancel(context.Background())

	go TestTick(ctx)

	time.Sleep(time.Second * 5)
	cancel()
	time.Sleep(time.Second * 5)

}
