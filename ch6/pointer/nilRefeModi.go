package main

func failToUpdate(g *int) {
	x := 10
	g = &x
}

func successToUpdate(g *int) {
	x := 10
	*g = x
}
