package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	ch := make(chan int)
	//ch := make(chan int,2)
	go square(&wg, ch)
	for i := 0; i < 10; i++ {
		ch <- i * 2
	}
	// 안닫으면 좀비 고루틴됨 (무한 대기)
	close(ch)
	fmt.Println("never print")
	wg.Wait()
}

func square(wg *sync.WaitGroup, ch chan int) {
	tick := time.Tick(time.Second)
	terminate := time.After(10 * time.Second)
	for {
		select {
		case <-tick:
			fmt.Println("tick")
		case <-terminate:
			fmt.Println("terminated ")
			wg.Done()
			return
		case n := <-ch:
			fmt.Println("square:", n*n)
			time.Sleep(time.Second)
		}
	}
}
