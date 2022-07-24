package main

import "fmt"

type person struct {
	name string
}

func main() {
	rizky := person{name: "rizky"}
	rizky.printName()
	rizky.changeName()
	rizky.printName()
}

func (p person) printName() {
	fmt.Println(p.name)
}

func (p *person) changeName() {
	p.name = "ganti"
}