package main

import (
	"fmt"
	"sort"
)

func main() {
	// Tạo slice chứa các số nguyên
	numbers := []int{5, 2, 9, 1, 5, 6}

	// Sắp xếp slice tăng dần
	sort.Ints(numbers)

	// In ra slice đã sắp xếp
	fmt.Println("numbers", numbers)