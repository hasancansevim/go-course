package main

import "fmt"

func main() {
	/*
		index := 1

		for index <= 10 {
			fmt.Println(index)
			index += 1
		}
	*/

	/*
		total := 0
		for index := 1; index <= 10; index++ {
			total += index
		}
		fmt.Println(total)
	*/

	/*
		var names = [3]string{"Mustafa", "Murat", "Coşkun"}

		for index := 0; index < len(names); index++ {
			fmt.Println(names)
		}
	*/

	/*
		for index := 0; index <= 10; index++ {
			if index == 3 {
				break
			}
			fmt.Println(index)
		}
	*/

	for index := 0; index <= 10; index++ {
		if index == 2 || index == 5 {
			continue
		}
		fmt.Println(index)
	}
}
