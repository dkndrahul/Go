package main

import "fmt"

type contactInfo struct {
	email   string
	zipcode int
}

type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// rahul :=  person{"Rahul", "Deekonda"} one way of delcaring
	// rahul := person{firstName: "Rahul", lastName: "Deekonda"} // other way of declaring with property name like key value pair

	// var rahul person

	// rahul.firstName = "Rahul"
	// rahul.lastName = "Deekonda"
	// fmt.Printf("%+v", rahul)
	// fmt.Println(rahul)

	rahul := person{
		firstName: "Rahul",
		lastName:  "Deekonda",
		contact: contactInfo{
			email:   "dkndrahul@gmail.com",
			zipcode: 503001,
		},
	}

	rahul.updateName("Raul")
	rahul.print()
}

func (pointerToPerson *person) updateName(nameString string) {
	(*pointerToPerson).firstName = nameString
}

func (p person) print() {
	fmt.Printf("%+v", p)
}
