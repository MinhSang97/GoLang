package model

import "fmt"

type Person struct {
	FirstName      string
	FullName       string
	Address        string
	IdentityNumber uint64
	ClassName      string
	Spouse         *Person
}

func (p Person) Sleep() {
	fmt.Println("Sleep 2 hours")

}

func (p Person) Eat() {
	fmt.Println("Eat 2 cakes")
}

func (p *Person) Marry(s Person) Person {
	p.Spouse = &s
	return *p
}

func (p *Person) GetSpouse() *Person {
	fmt.Println("debug", p.Spouse)
	return p
}
