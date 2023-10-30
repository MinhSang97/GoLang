package main

import (
	"fmt"
)

func main() {
	// Tạo slice chứa các số nguyên
	numbers := []int{5, 2, 9, 1, 4, 6}
	fmt.Println("Slice ban đầu", numbers)
	fmt.Println("-------------")
	// Thêm một phần tử vào đầu slice
	newElement := 0
	updatednumbers := append([]int{newElement}, numbers...)

	fmt.Println("Slice sau khi thêm phần tử:", updatednumbers)

}
