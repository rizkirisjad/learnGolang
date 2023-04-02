package main

import (
	"fmt"
)

func main() {
	counter := 0
	name := "Rizki"

	increment := func() {
		name = "Sasongko"
		fmt.Println("ok")
		counter ++
	}

	increment()
	increment()
	
	fmt.Println(counter)
	fmt.Println(name)
}

