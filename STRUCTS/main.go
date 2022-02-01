package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

func main() {
	fmt.Println("Starting....")
	me := person{firstName: "Alex", lastName: "Madurga"}
	me.firstName = "Alex"
	fmt.Println("Name: ", me.firstName)
}
