package main

import "fmt"

func main() {
	//willBePanic()
	useNewKeyWord()
}

// new 키워드는 포인터를 반환 단 비어있는 값은 제로값을 반환 nil 이아님

func useNewKeyWord() {
	var x = new(int)
	fmt.Println(x == nil)
	fmt.Println(*x)
}

// nil 값을 역참조시 패닉이 발생함
func willBePanic() {
	var x *int
	fmt.Println(x == nil)
	fmt.Println(*x)
}
