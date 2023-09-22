package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	account := &account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			for {
				depositAndWithdraw(account)
			}
			wg.Done()
		}()

	}
	wg.Wait()
}

func depositAndWithdraw(a *account) {
	if a.Balance < 0 {
		panic(fmt.Sprintf("balance should not be negative value:%d", a.Balance))
	}
	a.Balance += 1000
	time.Sleep(time.Millisecond)
	a.Balance -= 1000
}

type account struct {
	Balance int
}
