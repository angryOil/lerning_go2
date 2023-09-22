package main

import "fmt"

func main() {
	fir()
	sec()
	capTest()
	useMake()
	sliceNilNotNil()
	sliceSlicing()
	fullSlice()
	copySlice()
	copyWithArr()
}

func fir() {
	var x [3]int
	fmt.Println(x)
}

func sec() {
	var x = [3]int{1, 2: 2}
	fmt.Println(x)
}

func capTest() {
	var x []int
	fmt.Println(x, len(x), cap(x))
	x = append(x, 1)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 3)
	fmt.Println(x, len(x), cap(x))
}

func useMake() {
	x := make([]int, 5)
	fmt.Println(x, len(x), cap(x))
	x = append(x, 2)
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 1, 10)
	fmt.Println(y, len(y), cap(y))

	y = append(y, 2)
	fmt.Println(y, len(y), cap(y))
}

func sliceNilNotNil() {
	var nilData []int
	fmt.Println("nilData:", nilData)
	fmt.Println(nilData, len(nilData), cap(nilData))
	fmt.Println("isNil?", nilData == nil)
	var notNilData = []int{}
	fmt.Println("notNilData", notNilData)
	fmt.Println(notNilData, len(notNilData), cap(notNilData))
	fmt.Println("isNil?", nil == notNilData)
}

func sliceSlicing() {
	x := []int{1, 2, 3, 4}
	y := x[:2]
	z := x[1:]
	d := x[1:3]
	e := x[:]
	x = append(x, 100)
	y = append(y, 123)
	fmt.Println(len(x), cap(x))
	fmt.Println(len(y), cap(y))
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(d)
	fmt.Println(e)

}

func fullSlice() {
	x := []int{1, 2, 3, 4, 5}
	y := x[:2:2]
	z := x[2:4:4]
	z[1] = 99
	z = append(z, 123)
	fmt.Println(x, len(x), cap(x))
	fmt.Println(y, len(y), cap(y))
	fmt.Println(z, len(z), cap(z))

}

func copySlice() {
	x := []int{1, 2, 3, 4}
	y := make([]int, 4)
	copy(y, x)
	y[0] = 999
	fmt.Println(x)
	fmt.Println(y)

	a := []int{1, 2, 3, 4}
	b := make([]int, 2)
	c := make([]int, 2)
	copy(b, a)
	copy(c, a[3:])
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}

func copyWithArr() {
	x := []int{1, 2, 3, 4}
	y := [4]int{5, 6, 7, 8}
	z := make([]int, 8, 8)
	copy(z, x)
	copy(z[4:], y[:])
	fmt.Println(z)
}
