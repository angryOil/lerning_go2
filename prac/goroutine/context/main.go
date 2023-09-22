package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	//ctx, cancel := context.WithCancel(context.Background())
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	go printEverySecond(ctx)
	time.Sleep(5 * time.Second)
	cancel()
	wg.Wait()
}

func printEverySecond(ctx context.Context) {
	tick := time.Tick(time.Second)
	for {
		select {
		case <-ctx.Done():
			wg.Done()
			fmt.Println("done")
			return
		case <-tick:
			fmt.Println("tick ")

		}
	}
}
