package main

import "fmt"

type XEncoder struct {
}

type YEncoder struct {
}

type IEncoder interface {
	Encode(value string)
	Decode(value string)
}

func (xEncoder *XEncoder) Encode(value string) {
	fmt.Println("XEncoder ile encode edildi!")
}

func (xEncoder *XEncoder) Decode(value string) {
	fmt.Println("XEncoder ile decode edildi!")
}

func (yEncoder *YEncoder) Encode(value string) {
	fmt.Println("YEncoder ile encode edildi!")
}

func (yEncoder *YEncoder) Decode(value string) {
	fmt.Println("YEncoder ile decode edildi!")
}

func main() {
	var encoder IEncoder
	encoder = &YEncoder{}
	encoder.Encode("123456")
	encoder.Decode("xasd123gs53")
}
