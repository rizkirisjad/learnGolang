package main

import "fmt"

func main() {

	var values = [3]int{
		1,
		5,
		7,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(len(values))
}