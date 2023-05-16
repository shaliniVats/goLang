package main

import "fmt"

type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	colors := make(map[string]string)
	colors["white"] = "#ffffff"
	fmt.Println(colors)
	// var alex person
	// alex.firstName = "Alex"
	// alex.lastName = "Anderson"
	// alex.contact.email = "a_a@xyz.com"
	// alex.contact.zipCode = 123456

	// //alexPointer := &alex
	// alex.updateName("aloha")

	// fmt.Printf("%+v", alex)
}

func (PointerToPerson *person) updateName(newFirstname string) {
	(*PointerToPerson).firstName = newFirstname
}
