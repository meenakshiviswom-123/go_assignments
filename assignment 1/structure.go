package main

import "fmt"

type Address struct {
	Name    string
	city    string
	Pincode int
}

func main() {

	var a Address
	fmt.Println(a)

	a1 := Address{"Meenakshi Viswom", "Kerala", 689501}

	fmt.Println("Address1: ", a1)

	a2 := Address{Name: "Rajalekshmi", city: "Chennai",
		Pincode: 600001}

	fmt.Println("Address2: ", a2)

	a3 := Address{Name: "Praveen"}

	fmt.Println("Address3: ", a3)
}
