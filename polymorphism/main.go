package main

import "fmt"

type person struct {
	fname string
	lname string
}

type secretAgent struct {
	person
	licenceToKill bool
}
type human interface {
	speak()
}

func (p person) speak() {
	fmt.Println(p.fname, p.lname, `says, "Good Morning!"`)
}

func (sa secretAgent) speak() {
	fmt.Println(sa.fname, sa.lname, `says, "Shaken not stirred!"`)
}

func saySomething(h human) {
	h.speak()
}
func main() {
	p := person{
		fname: "Alice",
		lname: "D'souza",
	}

	sa := secretAgent{
		person{
			fname: "Vyomkesh",
			lname: "Bakshi",
		},
		true,
	}

	saySomething(p)
	saySomething(sa)
}
