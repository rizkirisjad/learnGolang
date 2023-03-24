package main

import "fmt"

func main() {
	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar Golang"
	book["author"] = "Rizki"
	fmt.Println(book)
	fmt.Println(len(book))
}