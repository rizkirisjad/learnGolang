package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpRizki NoKTP = "123456789"
	var marriedStatus Married = true
	fmt.Println(noKtpRizki)
	fmt.Println(marriedStatus)
}
