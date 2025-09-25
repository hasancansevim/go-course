package main

import "fmt"

func main() {
	/*
		var numbers = []int{1, 2, 3, 4, 5}

		for index, value := range numbers {
			fmt.Println(index, value)
		}

		for _, value := range numbers {
			fmt.Println(value)
		}
	*/

	/*
		var language string = "GoLang"

		for _, character := range language {
			fmt.Printf("Character: %c \n", character)
		}
	*/

	var names = map[string]int{
		"Mustafa": 10,
		"Murat":   20,
		"Coşkun":  30,
	}

	for key, value := range names {
		fmt.Printf("Key : %s , Value : %d \n", key, value)
	}
}
