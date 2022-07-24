package main

import "fmt"

type contactInfo struct {
	email string
	zipCode int
}

type person struct {
	firstName string
	lastName  string
	contactInfo
}

func main() {
	// rizky := person{firstName: "Rizky", lastName: "Abdulrachman"}
	// var rizky person

	// rizky.firstName = "rizky"
	// rizky.lastName = "abdulrachman"

	// fmt.Println(rizky)
	// fmt.Printf("%+v", rizky)

	jim := person {
		firstName: "Jim",
		lastName: "Party",
		contactInfo: contactInfo {
			email: "jim@gmail.com",
			zipCode: 12345,
		},
	}
	//  get memory address of variable jim value
	// jimPointer := &jim
	// jimPointer.updateName("jimmy")

	jim.updateName("berubah")
	jim.print()
	
}
// * pointer to person
func (pointerToPerson *person) updateName(newFirstName string){
	//  get value from memory address then change to newFirstName
	(*pointerToPerson).firstName = newFirstName
}

func (p person) print(){
	fmt.Printf("%+v", p)
}