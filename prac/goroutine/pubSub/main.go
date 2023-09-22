package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// 옵저버 패턴
var wg sync.WaitGroup

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	wg.Add(4)
	pub := newPublisher(ctx)
	sub1 := newSubscriber("A", ctx)
	sub2 := newSubscriber("B", ctx)

	go pub.update()
	sub1.subscribe(pub)
	sub2.subscribe(pub)

	go sub1.update()
	go sub2.update()

	go func() {
		tick := time.Tick(time.Second * 2)
		for {
			select {
			case <-tick:
				pub.publish("hello msg")
			case <-ctx.Done():
				wg.Done()
				return
			}
		}
	}()
	fmt.Scanln()
	cancel()
	wg.Wait()
}
