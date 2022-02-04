package main

import (
	"fmt"
)

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
	fmt.Println("Starting....")
	me := person{
		firstName: "Alex",
		lastName:  "Madurga",
		contact: contactInfo{
			email:   "amaduian@gnmail.com",
			zipCode: 28000,
		},
	}
	me.print()
	me.updateName("Pedro")
	me.print()

}

func (p person) print() {
	fmt.Printf("%+v", p)
	fmt.Println()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}
