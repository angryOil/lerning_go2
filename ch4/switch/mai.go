package main

import "fmt"

func main() {
	swi()
	swi2()
	swi3()
}

func swi3() {
	words := []string{"hello", "world", "joy", "warmOil", "헬로"}

	for _, word := range words {
		switch wordLen := len(word); wordLen {
		case 3:
			fmt.Println(word, "는 3글자입니다")
		default:
			fmt.Println("나머지 입니다")
		}
	}
}

func swi() {
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("it`s even", i)
		case i%3 == 0:
			fmt.Println(i)
		case i%7 == 0:
			fmt.Println(i, "종료")
			return
		}
	}
	fmt.Println("끝")
}

func swi2() {
loop:
	for i := 0; i < 10; i++ {
		switch {
		case i%2 == 0:
			fmt.Println("it`s even", i)
		case i%3 == 0:
			fmt.Println(i)
		case i%7 == 0:
			fmt.Println(i, "종료")
			break loop
		}
	}
	fmt.Println("끝")
}
