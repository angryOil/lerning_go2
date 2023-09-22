package main

import "fmt"

func main() {
	x := "hello"
	xp := &x
	fmt.Println(xp)
	fmt.Println(*xp)

	var n map[struct{}]struct{}
	fmt.Println(n)
	fmt.Println(n == nil)
	p := per{
		n: "n",
		i: 2,
	}
	chane := changePer(p)
	fmt.Println(p)
	fmt.Println(chane)
}

type per struct {
	n string
	i int
}

func changePer(p per) per {
	p.n = "modify"
	return p
}
