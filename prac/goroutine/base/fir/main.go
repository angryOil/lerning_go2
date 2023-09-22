package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	//printFunc()
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			sumAtoB(1, 1000000000)
		}()
	}
	wg.Wait()
}

func sumAtoB(a, b int) {
	sum := 0
	for i := a; i <= b; i++ {
		sum += i
	}
	fmt.Printf("%d 부터 %d 까지의 합은 %d 입니다 \n ", a, b, sum)
	wg.Done()
}

func printFunc() {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		printHan()
		wg.Done()
	}()
	go func() {
		printNum()
		wg.Done()
	}()
	wg.Wait()
}

func printHan() {
	hanguls := []rune{'가', '나', '다', '라', '마', '바', '사'}

	for _, v := range hanguls {
		time.Sleep(300 * time.Millisecond)
		fmt.Printf("%c", v)
	}
}

func printNum() {
	for i := 1; i <= 5; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%d", i)
	}
}
