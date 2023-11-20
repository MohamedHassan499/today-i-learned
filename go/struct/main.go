package main

import "fmt"

type contact struct {
	email   string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contact
}

func main() {
	alex := person{firstName: "Alex", lastName: "Anderson", contact: contact{email: "farag", zipCode: 12345}}
	// fmt.Println(alex)
	// fmt.Printf("%+v", alex)
	// fmt.Println(alex.contact.email)
	// fmt.Println(alex.contact.zipCode)
	alexP := &alex
	alexP.updateFirstName("Farag")
	alex.print()
}

// Pass by value
// func (p person) updateFirstName (firstName string) {
// 	p.firstName = firstName
// }

func (p *person) updateFirstName (firstName string) {
	(*p).firstName = firstName
}

func (p person) print() {
	fmt.Printf("%+v", p)
}