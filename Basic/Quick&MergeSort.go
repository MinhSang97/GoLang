package main

import (
	"fmt"
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

	// In ra slice ban đầu
	fmt.Println("Slice ban đầu:")
	fmt.Println(students)

	// Tạo một bản sao của slice để sắp xếp bằng Quick Sort
	quickSortStudents := make([]Student, len(students))
	copy(quickSortStudents, students)

	// Sắp xếp slice theo tên bằng thuật toán Quick Sort
	fmt.Println("Slice sau khi sắp xếp bằng Quick Sort:")
	quickSort(quickSortStudents)
	fmt.Println(quickSortStudents)

	// Tạo một bản sao của slice để sắp xếp bằng Merge Sort
	mergeSortStudents := make([]Student, len(students))
	copy(mergeSortStudents, students)

	// Sắp xếp slice theo tên bằng thuật toán Merge Sort
	fmt.Println("Slice sau khi sắp xếp bằng Merge Sort:")
	mergeSortStudents = mergeSort(mergeSortStudents)
	fmt.Println(mergeSortStudents)
}

// Đoạn mã Quick Sort
func quickSort(students []Student) {
	if len(students) < 2 {
		return
	}

	pivot := students[0]
	var left, right []Student

	for _, student := range students[1:] {
		if student.Name < pivot.Name {
			left = append(left, student)
		} else {
			right = append(right, student)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(students, left)
	students[len(left)] = pivot
	copy(students[len(left)+1:], right)
}

// Đoạn mã Merge Sort
func mergeSort(students []Student) []Student {
	if len(students) <= 1 {
		return students
	}

	mid := len(students) / 2
	left := mergeSort(students[:mid])
	right := mergeSort(students[mid:])

	return merge(left, right)
}

func merge(left, right []Student) []Student {
	result := make([]Student, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}
		if left[0].Name <= right[0].Name {
			result = append(result, left[0])
			left = left[1:]
		} else {
			result = append(result, right[0])
			right = right[1:]
		}
	}
	return result
}
