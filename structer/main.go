package main

import "fmt"

func main() {
	var customer1 = Customer{id: 1, name: "Mustafa Murat Coşkun", age: 30, address: Adress{city: "İstanbul", district: "Ataşehir"}}
	/*
		var customer2 = Customer{id: 2, name: "Hasan Can Sevim", age: 23}
		customer1.age = 31
		fmt.Println(customer1)
		fmt.Println(customer2)
	*/
	customer1.ToString()
	//changeName(&customer1)
	customer1.changeName("Hasan Can")
	customer1.ToString()
}

type Customer struct {
	id      int64
	name    string
	age     int
	address Adress
}

type Adress struct {
	city     string
	district string
}

func (customer Customer) ToString() {
	fmt.Printf("Name : %s , Age : %d\n", customer.name, customer.age)
}

/*
func changeName(customer *Customer) {
	customer.name = "Hasan Sevim"
}
*/

func (customer *Customer) changeName(newName string) {
	customer.name = newName
}
