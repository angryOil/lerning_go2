package main

import "fmt"

func main() {
	useForOk(countTo(10))
}

func useForOk(a <-chan int) {
	isRunning := true
	for isRunning {
		i, ok := <-a
		if i > 8 {
			isRunning = false
		}
		fmt.Println(i, ok)
	}
	result, ok := <-a
	fmt.Println(result, ok)
	select {
	case v, ok := <-a:
		fmt.Println(v, ok)
	default:
		fmt.Println("end")
	}
}

func justForLoop() {
	var a <-chan int = countTo(10)

	for i := range a {
		//_, ok := <-a
		if i > 8 {
			break
		}
		fmt.Println(i)
	}
	result, ok := <-a
	fmt.Println(result, ok)
	select {
	case v, ok := <-a:
		fmt.Println(v, ok)
	default:
		fmt.Println("end")
	}
}

func countTo(max int) <-chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i < max; i++ {
			ch <- i
		}
		close(ch)
	}()
	return ch
}
