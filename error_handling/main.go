package main

import (
	"fmt"
)

func main() {
	/*
		var pointer1 *int

		if pointer1 == nil {
			fmt.Println("Pointer herhangi bir yeri göstermiyor")
		}
	*/
	/*
		fileData, err := os.ReadFile("sample.txt")

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(fileData)
		}
	*/
	/*
		result, err := divide(10, 0)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	*/
	/*
		productService := ProductService{}
		err := productService.Add(Product{id: 1, name: "Ürün 1", price: -10})

		if err != nil {
			fmt.Println(err)
		}
	*/

	productService := ProductService{}
	err := productService.Add(Product{id: 1, name: "", price: -10})
	if err != nil {
		fmt.Println(err)
	}
}

/*
func divide(x int, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Payda 0 olamaz")
	}
	return x / y, nil
}
*/

type Product struct {
	id    int
	name  string
	price int
}

type ProductService struct {
}

func (productService *ProductService) Add(product Product) error {
	if len(product.name) == 0 {
		return &ValidationError{text: "Ürün İsmi Boş Olamaz", code: "400"}
	}
	if product.price < 0 {
		return &ValidationError{text: "Ürün Fiyatı 0 dan büyük olmalı", code: "400"}
	}
	fmt.Println("Ürün Eklendi")
	return nil
}

type ValidationError struct {
	code string
	text string
}

func (validationError ValidationError) Error() string {
	return fmt.Sprintf("Hata : %s , Kod : %s", validationError.text, validationError.code)
}
