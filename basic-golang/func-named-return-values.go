package main

import (
	"fmt"
)

func completeName() (firstName, lastName string, age int) {
	firstName = "Rizki"
	lastName = "Sasongko"
	age = 28

	return
}

func main() {
	fmt.Println(completeName())
}