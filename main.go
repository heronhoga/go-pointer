package main

import "fmt"

type Address struct {
	City,
	Province,
	Country string
}

func main() {
	address1 := Address{"Semarang", "Jawa Tengah", "Indonesia"}
	// address2 := address1 //copy value
	address2 := &address1 //copy value

	address2.City = "Yogyakarta"

	fmt.Println(address1)
	fmt.Println(address2)

	*address2 = Address{"Kota Gorontalo", "Gorontalo", "Indonesia"} //changed memory reference

	fmt.Println(address1)
	fmt.Println(address2)
}