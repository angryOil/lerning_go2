package main

import "fmt"

func main() {
	i := 1000
	failUpdate(&i)
	fmt.Println(i)
	update(&i)
	fmt.Println(i)
}

func failUpdate(i *int) {
	mod := 20
	i = &mod
}

func update(i *int) {
	*i = 22
}
