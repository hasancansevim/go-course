package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		myChannel := make(chan int, 2)
		myChannel <- 1
		myChannel <- 2
		fmt.Println("end of the main")
	*/
	/*
		myChannel := make(chan int, 2)
		myChannel <- 1
		myChannel <- 2
		// q yapısı:ilk giren ilk process ediliyor
		fmt.Println(<-myChannel) // channel'ın size ı 1 tane okuduğumuz için şu anda 1

		myChannel <- 3
		fmt.Println(<-myChannel) // size : 1 ,değer : 2
		fmt.Println(<-myChannel) // size :2  değer : 3

		fmt.Println(<-myChannel) // size 3 , hata
	*/
	messages := make(chan string, 2)
	go func() {
		data1 := <-messages
		fmt.Println("First : ", data1)
		data2 := <-messages
		fmt.Println("First : ", data2)
	}()

	messages <- "hello"
	messages <- "world"
	messages <- "world2"
	time.Sleep(time.Second)
	fmt.Println("end of the main")
}
