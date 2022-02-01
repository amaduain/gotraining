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
	me.firstName = "Alex"
	fmt.Println("Name: ", me.firstName)
	fmt.Printf("%+v", me)
}
