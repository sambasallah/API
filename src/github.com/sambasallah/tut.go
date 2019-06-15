package main

import (
	"fmt"
)

type Person struct {
	name,address string
	age int
}

func main() {
	person := &Person{name:"samba sallah",address:"serrekunda, the Gambia",age:30}

	fmt.Println(person.GetName(), person.GetAddress(), person.GetAge())
}

func (p *Person) GetName() (string) {
	return p.name;
}

func (p *Person) GetAddress() (string) {
	return p.address
}

func (p *Person) GetAge() (int) {
	return p.age
}