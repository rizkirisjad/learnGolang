package main

import "fmt"

func getHi(name string) string {
	return "Hi go" + name
}

func main() {
	result := getHi("Rizki")
	fmt.Println(result)
}