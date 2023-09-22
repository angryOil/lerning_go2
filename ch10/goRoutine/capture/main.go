package main

import "fmt"

func main() {
	buggedFun()
	fixedByShadowing()
	fixedByGoRoutineParam()
}

func fixedByShadowing() {
	fmt.Println("fixed by shadowing----")
	a := []int{1, 2, 3, 4, 5, 65}
	ch := make(chan int, len(a))
	for _, v := range a {
		v := v // type 생성 외부
		go func() {
			ch <- v * 2 // 타입 사용 안쪽
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func fixedByGoRoutineParam() {
	fmt.Println("fixed by goroutine param----")
	a := []int{1, 2, 3, 4, 5, 65}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func(val int) { // 타입 생성
			ch <- val * 2 // 타입 사용
		}(v)
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}

func buggedFun() {
	a := []int{1, 2, 3, 4, 5, 65}
	ch := make(chan int, len(a))
	for _, v := range a {
		go func() {
			ch <- v * 2
		}()
	}
	for i := 0; i < len(a); i++ {
		fmt.Println(<-ch)
	}
}
