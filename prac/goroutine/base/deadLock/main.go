package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	rand.Seed(time.Now().UnixNano())
	wg.Add(2)

	fork := &sync.Mutex{}
	spoon := &sync.Mutex{}
	go dingingProblem("A", fork, spoon, "포크", "수저")
	go dingingProblem("B", spoon, fork, "수저", "포")
	wg.Wait()
}

func dingingProblem(name string, first *sync.Mutex, second *sync.Mutex, firstStr string, secondStr string) {
	for i := 0; i < 100; i++ {
		fmt.Printf("%s 밥먹을 준비중 ... \n", name)
		first.Lock()
		fmt.Printf("%s %s 획득\n", name, firstStr)
		second.Lock()
		fmt.Printf("%s %s 획득\n", name, secondStr)

		fmt.Println("밥을 먹기 시작 ")
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

		second.Unlock()
		first.Unlock()
	}
	wg.Done()
}
