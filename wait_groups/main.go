package main

import (
	"fmt"
	"sync"
)

func main() {

	// projemizdeki join pointleri yönetmemizi sağlıyor
	// yani go routinleri / threadleri
	wg := sync.WaitGroup{}

	// go routine sayısı
	wg.Add(2)

	go func() {
		defer wg.Done() // sayacı 1 azaltır
		fmt.Println("f1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("f2")
	}()

	// go routinler bitene kadar alt satırlara geçmeyecek
	// sayacın 0 lanmasını bekler
	wg.Wait() // blocked

	fmt.Println("end of the main")
}
