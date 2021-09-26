package main

import "fmt"

func main() {
	type noKTP string
	type Married bool

	var noKtpMuhbit noKTP = "367102500000089"
	var marriedStatus Married = true

	fmt.Println(noKtpMuhbit)
	fmt.Println(marriedStatus)
}
