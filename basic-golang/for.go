package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Looping ke ", counter)
		counter++
	}

	for counter := 0; counter <= 9; counter++ {
		fmt.Println("Looping ke ", counter)
	}

	names := []string{"Rizki", "C", "Sasongko", "Malika", "Malaikha"}

	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	for i, name := range names {
		fmt.Println("index ke", i, "=", name)
	}

	person := make(map[string]string)
	person["name"] = "Rizki"
	person["job"] = "Software Engineer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}