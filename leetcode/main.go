package main

import "fmt"

func main() {
	fmt.Print("Nhập column title: ")
	var columnTitle string
	fmt.Scanln(&columnTitle)

	result := titleToNumber(columnTitle)
	fmt.Println(result)
}

func titleToNumber(columnTitle string) int {
	result := 0
	pow := 1

	for i := len(columnTitle) - 1; i >= 0; i-- {
		ch := columnTitle[i]
		digit := int(ch - 'A' + 1)

		result += digit * pow
		pow *= 26
	}

	return result
}
