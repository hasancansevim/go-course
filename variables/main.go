package main

import "fmt"

func main() {
	/*
		var productName string
		var quantity int
		var discount float32
		var isInStock bool

		productName = "Golang Dersleri"
		quantity = 10
		discount = 0.37
		isInStock = true
		fmt.Println(productName, reflect.TypeOf(productName))
		fmt.Println(quantity, reflect.TypeOf(quantity))
		fmt.Println(discount, reflect.TypeOf(discount))
		fmt.Println(isInStock, reflect.TypeOf(isInStock))
	*/

	/*
		var productName string = "Golang Dersleri"

		productName := "Golang Dersleri"
		fmt.Println(productName)
	*/

	var productName string
	var quantity int
	var discount float32
	var isInStock bool

	productName = "Golang Dersleri"
	quantity = 10
	discount = 0.37
	isInStock = true

	fmt.Printf("ProductName : %s , Quantity : %d , Discount : %f , IsInStock : %t\n", productName, quantity, discount, isInStock)
}
