package main

import (
	"fmt"
	"sort"
)

// Định nghĩa cấu trúc Student
type Student struct {
	Name string
}

func main() {
	// Tạo slice chứa các Student
	students := []Student{
		{"Alice"},
		{"Zob"},
		{"Eve"},
		{"Huy"},
		{"Charlie"},
		{"Hanh"},
	}
	// Khai báo biến đếm số lần so sánh
	comparisons := 0

	// In ra slice ban đầu
	fmt.Println("Slice ban đầu:")
	fmt.Println(students)

	// In ra slice đã sắp xếp
	fmt.Println("Slice sort lần đầu tiên:")
	fmt.Println(students)

	// Sắp xếp slice theo tên (theo thứ tự ABC)
	sort.Slice(students, func(i, j int) bool {
		return students[i].Name < students[j].Name
	})

	// Sắp xếp slice theo tên bằng thuật toán Bubble Sort
	for i := 0; i < len(students); i++ {
		for j := 0; j < len(students)-i-1; j++ {
			// Tăng biến đếm số lần so sánh
			comparisons++
			if students[j].Name > students[j+1].Name {
				students[j], students[j+1] = students[j+1], students[j]
			}
		}
	}
	fmt.Println("----------------")
	fmt.Printf("Số lần so sánh: %d\n", comparisons)
	fmt.Println("----------------")
	fmt.Println("Kết quả sau khi chạy sort cho slice bằng for và if")
	fmt.Println(students)
}
