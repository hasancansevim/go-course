package main

import "fmt"

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func main() {
	book1 := Book{desi: 10}

	fmt.Println(book1)
	fmt.Println(book1.calculateShippingCost())

	books := []Book{{desi: 10}, {desi: 20}}
	fmt.Println(calculateTotalShippingCost(books))
}

func (book *Book) calculateShippingCost() int {
	return 5 + 2*book.desi
}

func (electronic *Electronic) calculateShippingCost() int {
	return 10 + 3*electronic.desi
}

func calculateTotalShippingCost(books []Book) int {
	total := 0

	for _, book := range books {
		total += book.calculateShippingCost()
	}
	return total
}
