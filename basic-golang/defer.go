package main

import (
	"fmt"
)

func logging() {
	fmt.Println("selesai memanggil func")
}

func runApplication() {
	defer logging()
	fmt.Println("Run app")
}

func main() {
	runApplication()
}