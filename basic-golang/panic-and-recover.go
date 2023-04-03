package main

import (
	"fmt"
)

func endApp() {
	message := recover()
	if message != nil {

		fmt.Println(message)
	}
	fmt.Println("App selesai")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("App error")
	}
	
	fmt.Println("App berjalan")
}

func main() {
	runApp(true)
}