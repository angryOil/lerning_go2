package main

import (
	"fmt"
	"golang.org/x/exp/rand"
)

func main() {
	doGoTo()
	anonymousFun()
}

func anonymousFun() {
	nums := func() []int {
		var result []int
		for i := 0; i < 10; i++ {
			result = append(result, i*2)
		}
		return result
	}
	fmt.Println(nums())
}

func doGoTo() {
	a := rand.Intn(10)
	fmt.Println("a:", a)
	for a < 100 {
		if a%4 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Println("go to 중간입니다")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)

}
