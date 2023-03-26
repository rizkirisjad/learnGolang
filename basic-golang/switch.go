package main

import "fmt"

func main() {
	var name = "Rizki"

	switch name {
	case "Rizki":
		fmt.Println("Hi, Rizki!")
	case "Malika":
		fmt.Println("Hi, Malika!")
	default:
		fmt.Println("Boleh kenalan?")
	}

	switch length := len(name) ; length > 5 {
	case true:
		fmt.Println("Terlalu Panjang!")
	case false:
		fmt.Println("Okee pas!")
	}
}