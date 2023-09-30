package main

import (
	"fmt"
)

func main() {
	loop()
	strLoop()
	loop2()
}

func loop2() {
	evenVals := []int{2, 4, 6, 8, 10}
	for i, v := range evenVals[1:] {
		fmt.Print(i+1, v, "\t")
	}
}

func strLoop() {
	str := "안녕하세요?abc"
	by := []byte(str)
	fmt.Println(by)
	for _, char := range by {
		fmt.Print(string(char))
	}
	fmt.Println()
	bToStr := string(by)
	fmt.Println(bToStr)
	for _, char := range bToStr {
		fmt.Print(string(char))
		fmt.Print(char, "\t")
	}
	fmt.Println()
}

func loop() {
	for i := 0; i < 10; i++ {
		if i < 5 {
			return
		}
		fmt.Print(i)
	}
}
