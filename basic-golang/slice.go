package main

import "fmt"

func main() {
	var mobileGame = []string{"PUBG", "LOL", "Apex Legends",}

	mobileGame = append(mobileGame, "Mobile Legends")

	fmt.Println(mobileGame)

	for _, mg := range mobileGame {
		fmt.Println(mg)
	}
}