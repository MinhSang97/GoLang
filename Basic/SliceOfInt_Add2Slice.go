package main

import "fmt"

func main() {
	// Tạo một slice ban đầu
	aslice := []int{1, 2, 3, 4, 5}
	fmt.Println("aSlice ban đầu:", aslice)
	fmt.Println("-------------")

	bslice := []int{6, 7, 8, 9, 10}
	fmt.Println("bSlice ban đầu:", bslice)
	fmt.Println("-------------")

	// add thêm phần tử cuối Slice
	newslice := append(aslice, bslice...)

	fmt.Println("Slice sau khi add bslice vào đuôi aslice:", newslice)
}
