package main

import (
	"fmt"
	"math/rand"
)

func main() {
	useLabel()
	useCopy()
	useSwitch()
	blankSwitch()
	doYouKnowGoTO()
	rightUseGoto()
}

func useCopy() {
	evenVals := []int{1, 2, 3, 4, 5}
	doubleVals := []int{}
	for _, v := range evenVals {
		doubleVals = append(doubleVals, v*2)
	}
	fmt.Print(doubleVals)
}

func useLabel() {
	samples := []string{"hello", "world???한글"}
outer:
	for _, sample := range samples {
		for _, v := range sample {
			if v == 'l' {
				continue outer
			}
			fmt.Println(v, string(v))
		}
	}
}

func useSwitch() {
	a := 'a'

	switch size := byte(a); size {
	case 'a':
		fmt.Println("hello?")
		fallthrough
	default:
		fmt.Println("default")
	}
}

func blankSwitch() {
	words := []string{"hi", "hello", "world2"}
	for _, word := range words {
		switch wordLen := len(word); {
		case wordLen > 3:
			fmt.Println(word, " is wordLen over 3")
		default:
			fmt.Println(word, " is wordLen not over 3")
		}
	}
	switch {
	case len(words) > 1:
		fmt.Println("words len is over 1")
	default:
		fmt.Println("words len is over 1")
	}
}

func doYouKnowGoTO() {
	a := 10
	b := 20
	goto skip
skip:
	c := 30
	fmt.Println(a, b, c)

}

func rightUseGoto() {
	a := rand.Intn(10)
	for a < 100 {
		if a%5 == 0 {
			goto done
		}
		a = a*2 + 1
	}
	fmt.Print("do something when the loop completes normally")
done:
	fmt.Println("do complicated stuff no matter why we left the loop")
	fmt.Println(a)
}
