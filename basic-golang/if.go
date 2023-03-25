package main

import "fmt"

func main() {
	var name = "Rizki"
	name = "Malika"
	name = "Sasongko"

	if name == "Rizki" {
		fmt.Println("Hello Rizki!")
	} else if name == "Sasongko" {
		fmt.Println("Hi, S!")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name) ; length > 5 {
		fmt.Println("Terlalu panjang!")
	} else {
		fmt.Println("Okee pas!")
	}
}