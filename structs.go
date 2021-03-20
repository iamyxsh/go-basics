package main

import "fmt"

type contactInfo struct {
	email string
	pincode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	yash := person{
		firstName: "Yash",
		lastName: "Sharma",
		contactInfo : contactInfo {
			email: "yash@gmail.com",
			pincode: 700307,
		},
	}

	fmt.Printf("%+v", yash)
}

