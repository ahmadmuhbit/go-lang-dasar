package main

import "fmt"

func main() {

	var name = "Eko"

	if name == "Muhbit" {
		fmt.Println("Hello Muhbit")
	} else if name == "Joko" {
		fmt.Println("Hello Joko")
	} else if name == "Budi" {
		fmt.Println("Hello Budi")
	} else {
		fmt.Println("Hi kenalan dong")
	}

	// If dengan short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}
