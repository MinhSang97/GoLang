// getFullName
package main

import (
	"fmt"
)

type student struct {
	FirstName string
	LastName  string
	Age       int
	Grade     float32
	ClassName string
}

func main() {

	sang := student{
		FirstName: "Sang",
		LastName:  "Nguyen",
		Age:       26,
		Grade:     10,
		ClassName: "12A4",
	}
	fmt.Println(getFullName(sang.FirstName, sang.LastName, sang.ClassName, sang))
	sang.setName("Huong", "Quynh")
	fmt.Println(sang)
	fmt.Println("Receiver func get full name of Student: ", sang.ReceiverGetFullName())
	fmt.Println("Receiver func get first name of Student: ", sang.ReceiverGetFirstName())

}

func getFullName(a, b, c string, d student) string {
	return d.FirstName + " " + d.LastName + " " + d.ClassName

}

func (c *student) setName(a, b string) {
	*&c.FirstName = a
	*&c.LastName = b
}

func (b student) ReceiverGetFullName() string {
	return b.FirstName + " " + b.LastName

}
func (b student) ReceiverGetFirstName() string {
	return b.FirstName

}
