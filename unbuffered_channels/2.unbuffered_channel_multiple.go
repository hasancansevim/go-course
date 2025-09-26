package main

import (
	"fmt"
	"time"
)

func main() {
	myChannel := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			myChannel <- i
			fmt.Println("Sent data : ", i)
			time.Sleep(time.Second)
		}
		close(myChannel)
	}()

	for {
		data, isOpen := <-myChannel
		if isOpen == false {
			break
		}
		fmt.Print("Recivered Data : ", data)
	}
}
