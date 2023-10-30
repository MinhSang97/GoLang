package main

import (
	"fmt"
	"time"
)

func main() {
	firstName := "Manh"
	lastName := "Nguyen"
	fullName := firstName + "+" + lastName
	i := 0

	msg := "hello world 123"
	fmt.Println("hello world")
	fmt.Println(msg)
	fmt.Println(firstName, "+", lastName)
	fmt.Println(fullName)
	fullName = firstName + "-" + firstName
	fmt.Println(fullName)

	for i := 0; i < 10; i++ {
		fmt.Println(i+1, fullName)
	}

	for {
		fmt.Println(i+1, fullName)
		time.Sleep(time.Millisecond * 100)
		i = i + 1
		if i >= 10 {
			break

		}

	}

}
// go run Day1_290823.go  