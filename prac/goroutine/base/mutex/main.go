package main

import (
	"fmt"
	"sync"
	"time"
)

// 상호 배제
// mutual exclusion
var mutex sync.Mutex
var wg sync.WaitGroup

func main() {
	wg.Add(10)
	acc := &account{10}
	for i := 0; i < 10; i++ {
		for {
			go func() {
				depositAndWithdraw(acc)
			}()
		}
		wg.Done()
	}
	wg.Wait()
}

func depositAndWithdraw(a *account) {
	mutex.Lock()
	defer mutex.Unlock()
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
