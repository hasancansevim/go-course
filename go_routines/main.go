package main

import (
	"fmt"
	//"time"
)

func main() {
	go f1()
	go f2()
	fmt.Println("end of the  main")
	//time.Sleep(2 * time.Second)
}

func f1() {
	fmt.Println("f1")
	// time.Sleep(2 * time.Second)
}

func f2() {
	fmt.Println("f2")
}
