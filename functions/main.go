package main

import "fmt"

func main() {
	/*
			//total := add(10, 20)
			total, differance := calculation(10, 20)
			fmt.Println(total)
			fmt.Println(differance)

		var numbers = []int{10, 20, 30, 40, 50}

		total := sum(numbers)
		fmt.Println(total)

	*/
	fmt.Println(sum(3, 4, 5, 6))
}

/*
func add(x int, y int) int {
	// fmt.Println(x + y)
	return x + y
}

func calculation(x int, y int) (int, int) {
	return x + y, x - y
}

func sum(numbers []int) int {
	total := 0

	for _, value := range numbers {
		total += value
	}

	return total
}
*/

func sum(numbers ...int) int {
	sum := 0
	for _, value := range numbers {
		sum += value
	}
	return sum
}
