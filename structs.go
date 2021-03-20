package main

import "fmt"

type contactInfo struct { 
	email string
	pincode int
}

type person struct { // defining a perosn (object) type struct
	firstName string  // name-of-the-field type
	lastName string
	contactInfo // embedding diff struct in another struct
	//same as writting contactInfo contactInfo
}

func main() {
	yash := person{                       // defining yash as a new person
		firstName: "Yash",
		lastName: "Sharma",
		contactInfo : contactInfo {
			email: "yash@gmail.com",
			pincode: 700307,
		},
	}

	pointerToYash := &yash  // getting the address of the variable yash 
							// & gets you the address of the variable in ram
							// for eg. ->  1001110

	pointerToYash.updateName("Yxsh", "Shxrma") // calling the function using the pointer

	yash.print()
}

func(p person) print(){       // recieving p as a person
		fmt.Printf("%+v", p)
}

func(pointerToPerson *person) updateName( firstName string, lastName string){  // recieving address (pointer) which is of type pointer to person
		(*pointerToPerson).firstName = firstName  // updating the person who is at the pointer pointerToPerson
		(*pointerToPerson).lastName = lastName // * returns the value stored at that pointer
}



