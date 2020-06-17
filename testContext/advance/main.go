package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type Server struct {
	wg sync.WaitGroup
}

func TTT(ctx context.Context, wg *sync.WaitGroup) {
	<-ctx.Done()
	fmt.Println("get done")
	fmt.Println("get done222")
	wg.Done()
	return
}

func (s *Server) StartSyncBinlogService(ctx context.Context) {

	s.wg.Add(1)
	go func(ctx context.Context) {
		fmt.Println("add group wait")
		TTT(ctx, &s.wg) // 必须是指针传入
		//syncBinlog(ctx, inst, sysCfg)
		fmt.Println("wait group done")
	}(ctx)
}

func (s *Server) TestCancel(cancel context.CancelFunc) {
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		fmt.Printf("get signal %v\n", sig)
		cancel()
		fmt.Println("http service shut down")
	}()
	s.wg.Wait()

}

func main() {

	ctx, cancel := context.WithCancel(context.Background())

	s := new(Server)
	s.StartSyncBinlogService(ctx)
	s.TestCancel(cancel)

}
