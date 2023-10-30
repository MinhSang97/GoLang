package main

import (
	"app/model"
	"fmt"
)

func main() {
	a := model.Person{
		FirstName: "Sang",
		FullName:  "Nguyen Minh Sang",
		Address:   "Go Vap",
	}

	b := model.Person{
		FirstName: "Manh",
		FullName:  "Nguyen Duc Manh",
		Address:   "Go Vap",
	}

	student := model.Student{
		Person:    model.Person{FullName: "Nguyen Minh Sang"},
		FirstName: "Manh",
		LastName:  "Nguyen Duc Manh",
		Age:       26,
	}

	a.Marry(b)
	a.Sleep()
	a.GetSpouse()
	fmt.Println(a.GetSpouse())
	student.Sleep()
	student.Marry(a)
	fmt.Println(student.GetSpouse())
	fmt.Println(student)
	fmt.Printf("%+v", student)
	fmt.Println(student.FullName)
}
