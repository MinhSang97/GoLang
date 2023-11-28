package main

import (
	"fmt"
	"strings"
)

func main() {

	string := "     hello huu hong   "

	string = strings.ReplaceAll(string, " ", "")

	fmt.Println(string)
}
