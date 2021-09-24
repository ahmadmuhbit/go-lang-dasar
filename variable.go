package main

import "fmt"

func main() {
	var name string

	name = "Ahmad Muhbit"
	fmt.Println(name)

	name = "Muhbit Ahmad"
	fmt.Println(name)

	// Tidak menuliskan tipe datanya
	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	// Tidak menggunakan var
	country := "Indonesia"
	fmt.Println(country)

	country = "America"
	fmt.Println(country)

	// Deklarasi multiple variable
	var (
		firstName = "Ahmad"
		lastName  = "Muhbit"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
