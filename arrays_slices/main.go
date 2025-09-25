package main

import "fmt"

func main() {
	// Fixed
	/*
		var names [3]string
		names[0] = "Mustafa"
		names[1] = "Murat"
		names[2] = "Coşkun"
		fmt.Println(names)
	*/

	/*
		var names = [3]string{"Hasan", "Can", "Sevim"}
		names[0] = "Mustafa"

		fmt.Println(names)
		fmt.Println(names[0:1])
	*/

	// Slices

	var names = []string{"ahmet", "mehmet", "tolga"}
	names = append(names, "elif")
	fmt.Println(names)
}
