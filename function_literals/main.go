package main

import "fmt"

func main() {
	/*
		func() {
			fmt.Println("f1")
		}() // tanımladığımız yerde fonksiyonu çağırıyoruz

		func(x int, y int) {
			fmt.Println("x + y = ", x+y)
		}(3, 5)

		add := func(x int, y int) int {
			return x + y
		}
		fmt.Println(add(3, 2))
	*/
	add := func(x int, y int) int {
		return x + y
	}
	multiply := func(x int, y int) int {
		return x * y
	}
	a, b := calculator(3, 5, add, multiply)
	fmt.Println(a, b)
}

func calculator(x int, y int, f1 func(int, int) int, f2 func(int, int) int) (int, int) {
	return f1(x, y), f2(x, y)
}
