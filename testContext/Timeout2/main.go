package main

import (
	"context"
	"fmt"
	"time"
)

func TimeOutFunc(ctx context.Context) error {

	select {
	case <-ctx.Done():
		fmt.Println("timeout .......")
		return ctx.Err()
	}

}

func main() {

	timeoutCtx, _ := context.WithTimeout(context.Background(), time.Second*1)
	for {
		err := TimeOutFunc(timeoutCtx)
		if err == context.DeadlineExceeded {
			timeoutCtx, _ = context.WithTimeout(context.Background(), time.Second*1)

		}
	}

}
