package main

import "fmt"

func main() {
	tenAdder := adder{start: 10}
	tenAdderTo := adder.addTo
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdder.addTo(10))
	fmt.Println(tenAdderTo(tenAdder, 12))
	var a int = 200
	num(a).name()
}

type num int

func (n num) name() string {
	return fmt.Sprintf("my name is %d", n)
}

type adder struct {
	start num
}

func (a adder) addTo(val int) int {
	a.start += num(val)
	return int(a.start)
}
