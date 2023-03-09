package main

import "fmt"

func main() {
    i := 21
		j := true
		k := false
		ch := 'Я'
		num := 15
		kfloat := 123.456

    fmt.Println(i)
		fmt.Printf("%T\n", i)
		fmt.Printf("%%\n")
    fmt.Println(j)
		fmt.Println(k)
		fmt.Printf("%b\n", ch)
		fmt.Printf("%d\n", 21)
		fmt.Printf("%o\n", 025)
		fmt.Printf("%x\n", num)
    fmt.Printf("%X\n", num)
		fmt.Printf("%U\n", 'Я')
		fmt.Printf("%.6f\n", kfloat)
		fmt.Printf("%e\n", kfloat)
}
