package main

import "fmt"

func main() {
	/*
		var a int
		a = 10

		var p *int
		p = &a

		*p = 20
		fmt.Println(a)
		fmt.Println(p)
		fmt.Println(*p)
	*/
	var a = 10
	fmt.Println(a)
	add12pointer(&a)
	fmt.Println(a)
}

func add12pointer(x *int) {
	*x = *x + 12
}
