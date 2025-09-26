package main

import "fmt"

func main() {
	/*
		//go routine asenkron çalışacağı için message oluşmayacak
		message := go func () string {
			return "Hello"
		}()
	*/

	/*
		// go routine den bir değişkene değer atamak için channelları kullanıyoruz
		myChannel := make(chan string)

		go func() {
			myChannel <- "Hello World"
		}()

		// go routine myChannel'a bilgi basana kadar burası blocklanıyor
		message, isOpen := <-myChannel

		fmt.Println(message, isOpen)
	*/

	/*
		myChannel := make(chan string)

		go func() {
			message := <-myChannel
			fmt.Println(message)
		}()

		go func() {
			myChannel <- "Hello World"
		}()

		fmt.Println("end of the main")
	*/

	myChannel := make(chan string)
	done := make(chan bool)

	go func() {
		message := <-myChannel
		fmt.Println(message)
		done <- true
	}()

	go func() {
		myChannel <- "Hello World"
	}()

	<-done
	fmt.Print("end of the main")
}
