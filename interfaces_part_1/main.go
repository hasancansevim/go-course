package main

import "fmt"

type IShippable interface {
	ShippingCost() int
}

type Book struct {
	desi int
}

type Electronic struct {
	desi int
}

func main() {
	/*
		var product IShippable
			product = &Book{desi: 10}

			fmt.Println(product.ShippingCost())

			product = &Electronic{desi: 20}
			fmt.Println(product.ShippingCost())
	*/
	var products []IShippable = []IShippable{
		&Book{desi: 10},
		&Book{desi: 20},
		&Electronic{desi: 15},
	}

	fmt.Println(calculateTotalShippingCost(products))

}

func (book *Book) ShippingCost() int {
	return 5 + 2*book.desi
}

func (electronic *Electronic) ShippingCost() int {
	return 10 + 3*electronic.desi
}

/*
func calculateTotalShippingCost(books []Book) int {
	total := 0

	for _, book := range books {
		total += book.ShippingCost()
	}
	return total
}
*/

func calculateTotalShippingCost(products []IShippable) int {
	total := 0

	for _, product := range products {
		total += product.ShippingCost()
	}
	return total
}
