package main

import (
	"fmt"
)

type Student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {
	Manh := Student{
		FirstName: "Manh",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     9,
		ClassName: "HHM",
	}
	fmt.Println(Manh)

	SetName("manh1", "le", &Manh)
	fmt.Println(Manh)

	FullName := GetFullName3(Manh)
	fmt.Println(FullName + " 1")
}

func GetFullName3(student Student) string {
	return "Hello! I am " + student.FirstName + " " + student.LastName
}

func SetName(firstname, lastname string, student *Student) {
	student.FirstName = firstname
	student.LastName = lastname
}
