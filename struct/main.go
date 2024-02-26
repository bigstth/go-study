package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	alex := person{firstName: "123", lastName: "Test", contactInfo: contactInfo{email: "alex@mail.com", zipCode: 12123}}
	// alexPointer := &alex
	// alexPointer.updateName("111", "222")
	alex.updateName("111", "222") //shortcut
	alex.print()
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (pointerToPerson *person) updateName(firstName string, lastName string) {
	(*pointerToPerson).firstName = firstName
	(*pointerToPerson).lastName = lastName
}
