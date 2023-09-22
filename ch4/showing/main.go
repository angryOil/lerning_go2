package main

import (
	"fmt"
	"math/rand"
)

func main() {
	useVarShadowing()
	failShadow()
	badShadowing()
	shadowingUniverse()
	resolvedInIf()
}

func useVarShadowing() {
	fmt.Println("25.1 var shadowing")
	var x = 10
	y := 20
	if x > 5 {
		fmt.Print(x, ",")
		fmt.Println(y)
		var x = 5
		y := 5
		fmt.Print(x, ",")
		fmt.Println(y)
	}
	fmt.Println(x, ",", y)
}

func failShadow() {
	fmt.Println("fail to  showing")
	var x = 10
	y := 20
	if x > 1 {
		fmt.Println(x, ",", y)
		x = 5
		y = 5
		fmt.Println(x, ",", y)
	}
	fmt.Println(x, ",", y)
}

func badShadowing() {
	x := 10
	fmt.Println(x)
	fmt := 102
	println(fmt)
}

func shadowingUniverse() {
	fmt.Println(true)
	true := 22
	fmt.Println(true)
	string := 223
	fmt.Println(string)
}

func resolvedInIf() {
	fmt.Println(rand.Intn(20000))
	if n := rand.Intn(100); n == 0 {
		fmt.Println("0 is zero")
	} else {
		fmt.Println(n)
		fmt.Println("another num is :", n)
	}
}
