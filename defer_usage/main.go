package main

import "fmt"

// defer : ertelemek
// fonksiyonda nerde kullandıysak o statement ertelenecek ve fonksiyon bittiğinde çalışacak
func main() {
	/*
		defer fmt.Println("Hello")
		fmt.Println("World")
	*/
	/*
		defer fmt.Println("1. defer")  // 4
		defer fmt.Println("2. defer")  // 3
		defer fmt.Println("3. defer")  // 2
		fmt.Println("Main Fonksiyonu") // 1
	*/

	/*
		var condition = true
		defer fmt.Println("1. defer")
		if condition {
			return
		}
		defer fmt.Println("2. defer")
	*/
	/*
		x := 10
		y := 20

		defer fmt.Println("x:", x) // kullanıldığı andaki değerini esas alır
		defer fmt.Println("y:", y)
		x = 100
		y = 200

		fmt.Println("x:", x)
		fmt.Println("y:", y)
	*/

	condition := true
	defer cleanUp() // uygulama panic atmasına rağmen çalışacak
	if condition {
		panic("An error occurred") // uygulama panic'i görünce direkt olarak sonlanacak
	}

}

func cleanUp() {
	fmt.Println("Cleanup worked ...")
}
