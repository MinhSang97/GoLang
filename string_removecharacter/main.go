package main

import (
	"fmt"
	"strconv"
)

func main() {

	var input string
	fmt.Scanln(&input)

	number, _ := strconv.Atoi(input)

	s := "hello world"
	runes := []rune(s)

	if number >= 0 && number < len(runes) {
		runes = append(runes[:number], runes[number+1:]...)
	}

	fmt.Println(string(runes))

}
