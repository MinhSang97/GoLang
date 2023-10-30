package main

import (
	"fmt"
)

func main() {
	// Tạo slice chứa các số nguyên
	numbers := []int{5, 2, 9, 1, 4, 6}
	fmt.Println("Slice ban đầu", numbers)
	fmt.Println("-------------")
	// add thêm phần tử cuối Slice
	numbers = append(numbers, 11)
	fmt.Println("Slice sau khi thêm phần tử cuối", numbers)

}
