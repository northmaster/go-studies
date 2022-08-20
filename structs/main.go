package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName string
	contactInfo
}

func main() {
	jim := person{
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo{
			email: "jim@gmail.com",
			zipCode: 99999,
		},
	}

	
	jim.updateName("Charlie")
	jim.print()
}

func (p *person) updateName(newFirstName string) {
	p.firstName = newFirstName
}

func (p person) print() {
	fmt.Println(p)
}