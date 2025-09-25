package main

import "fmt"

func main() {
	// map [key]value
	/*
		var names map[string]int

		names = make(map[string]int, 0)
		names["Mustafa"] = 1
		names["Murat"] = 2
		names["Coşkun"] = 3

		fmt.Println(names)
		fmt.Println(names["Mustafa"])
		fmt.Println(names["Melike"])
	*/

	names := map[string]int{
		"Hasan": 1,
		"Can":   2,
		"Sevim": 3,
	}
	delete(names, "Can")
	fmt.Println(names)
}
