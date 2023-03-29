package main

import (
	"fmt"
)

func fullName() (string, string) {
	return "Rizki", "C"
}

func main() {
	firstName, lastName := fullName()
	fmt.Println(firstName, lastName)	
}