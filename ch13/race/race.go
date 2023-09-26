package race

import (
	"sync"
)

func getCounter() int {
	var counter int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(num int) {
			for i := 0; i < 1000; i++ {
				mu.Lock()
				counter++
				mu.Unlock()
			}
			wg.Done()
		}(i)
	}
	wg.Wait()
	return counter
}
