package main

import (
	"context"
	"fmt"
)

type Product struct {
	id   int64
	name string
}

var productChannel = make(chan Product)

func main() {
	/*
		go getProductDetailsFromExtarnalService(10)

		select {
		case productDetail := <-productChannel:
			fmt.Println("Ürün Detayları Getirildi ", productDetail)
		}

		fmt.Println("end of the channel")
	*/
	/*
		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		go getProductDetailsFromExtarnalService(10)

		select {
		case productDetail := <-productChannel:
			fmt.Println("Ürün Detayları Getirildi ", productDetail)
		case <-ctx.Done():
			fmt.Println("Timeout occurred,context cancelled")
		}

		fmt.Println("end of the channel")
	*/
	ctx := context.Background()
	ctx = context.WithValue(ctx, "correlation-id", "abc123")
	F1(ctx)
}

/*
func getProductDetailsFromExtarnalService(productId int64) {
	time.Sleep(10 * time.Second)
	productChannel <- Product{
		id:   productId,
		name: "Telefon",
	}
}
*/

func F1(ctx context.Context) {
	fmt.Println("F1", ctx.Value("correlation-id"))
	F2(ctx)
}
func F2(ctx context.Context) {
	fmt.Println("F2", ctx.Value("correlation-id"))
	F3(ctx)
}
func F3(ctx context.Context) {
	fmt.Println("F3", ctx.Value("correlation-id"))
}
