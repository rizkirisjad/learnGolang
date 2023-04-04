package main

import (
	"fmt"
)

type Customer struct {
	Name, Address string
	Age int
}

func (customer Customer) sayHelo(name string) {
fmt.Println("Hello", name, "My name is", customer.Name)
}

func main() {
	var risjad Customer
	risjad.Name = "Rizki"
	risjad.Address = "Indonesia"
	risjad.Age = 28

	fmt.Println(risjad)

	risjad.sayHelo("Joko")

	malika := Customer {
		Name: "Malika",
		Address: "Indonesia",
		Age: 25,
	}

	fmt.Println(malika)
}