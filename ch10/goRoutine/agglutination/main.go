package main

import "fmt"

func main() {
	doUKnowSelect()
}

func deadRock() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
	}()
	v := 2
	ch2 <- v
	v2 := <-ch1
	fmt.Print(v, v2)
}

func doUKnowSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		fmt.Println("go gogogo")
		v := 1
		ch1 <- v
		fmt.Println("go?")
		v2 := <-ch2
		fmt.Println(v, v2) // ?? << 무시됨
	}()
	v := 3
	var v2 int
	select {
	case ch2 <- v:
		fmt.Println("v??", ch2)
	case v2 = <-ch1:
		fmt.Println("ch1?", ch1)
	}
	fmt.Println(v, v2)
}
