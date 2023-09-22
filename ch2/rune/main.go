package main

import "fmt"

func main() {
	changeType()
}

func changeType() {
	var x int = 10
	var y float64 = 30.2
	var z float64 = float64(x) + y
	var d int = int(z)
	d, a := 223, 22
	fmt.Println(a)
	fmt.Println(z)
	fmt.Println(d)
}
