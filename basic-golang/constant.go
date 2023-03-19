package main

import "fmt"

func main() {
	const firstName string = "Rizki C"
	const lastName = "Sasongko"
	const value = 99999

	fmt.Println(firstName)
	fmt.Println(lastName)
	fmt.Println(value)

	const (
		friendName = "Messi"
		club       = "Psg"
		score      = 1000
	)

	fmt.Println(friendName)
	fmt.Println(club)
	fmt.Println(score)

}
