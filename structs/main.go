package main

import "fmt"

type contactInfo struct {
	email string
	zip   int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	aakash := person{
		firstName: "Aakash",
		lastName:  "Singhal",
		contact: contactInfo{
			email: "aakash@gmail.com",
			zip:   110089,
		},
	}
	fmt.Printf("%+v", aakash)
}
