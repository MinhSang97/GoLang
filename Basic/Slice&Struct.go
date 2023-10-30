package main

import (
	"fmt"
)

type Students struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {

	ten := "Sang"
	ho := "Nguyen"
	msg := getfullname(ten, ho)
	fmt.Println(msg)
	msg1 := getfullname2(ten, ho)
	fmt.Println(msg1)

	// Khai báo slice
	days := []string{"Monday", "Tuesday", "Wednesday"}
	msg2 := getDays(days)
	fmt.Println(msg2)

	// Struct
	Manh := Students{
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
	fmt.Println(FullName + " 'kết quả BTVN'")
}

func getfullname(firstname, lastname string) string {
	return "Hello! I am " + firstname + " " + lastname
}

// Cách 2
func getfullname2(firstname, lastname string) string {
	msg := fmt.Sprintf("Hello world! I am %s %s %d", firstname, lastname, 2023)
	return msg
}

func GetFullName3(student Students) string {
	return "Hello! I am " + student.FirstName + " " + student.LastName
}

func SetName(firstname, lastname string, student *Students) {
	student.FirstName = firstname
	student.LastName = lastname
}

func getDays(day []string) string {
	return day[0] + "," + day[1] + "," + day[2]
}
