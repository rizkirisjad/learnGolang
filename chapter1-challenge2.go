package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	// Looping untuk nilai i
	for i := 0; i < 5; i++ {
		fmt.Printf("nilai i = %d\n", i)
	}

	// Looping untuk nilai j
	for j := 0; j < 5; j++ {
		fmt.Printf("nilai j = %d\n", j)
	}

	// Slice dari sebuah string
	str := "CAШABO"
	slice := []byte(str)

	// Looping untuk mendapatkan karakter dan byte position
	pos := 0
	for len(slice) > 0 {
		char, size := utf8.DecodeRune(slice)
		fmt.Printf("character %c starts at byte position %d\n", char, pos)
		slice = slice[size:]
		pos += size
	}

	// Looping untuk nilai j lagi
	for j := 6; j < 11; j++ {
		fmt.Printf("nilai j = %d\n", j)
	}
}
