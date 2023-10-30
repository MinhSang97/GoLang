package main

import "fmt"

func main() {
	// Tạo một slice ban đầu
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice ban đầu:", slice)

	// Nhập giá trị n từ người dùng
	var n int
	fmt.Print("Nhập giá trị n: ")
	fmt.Scanln(&n)

	// Kiểm tra xem n có hợp lệ không
	if n < 0 || n >= len(slice) {
		fmt.Println("n không hợp lệ.")
		return
	}

	// Tách slice thành hai phần
	firstPart := slice[:n]
	secondPart := slice[n+1:]

	// Tạo slice mới bằng cách kết hợp phần đầu và phần cuối
	updatedSlice := append(firstPart, secondPart...)

	fmt.Println("Slice sau khi xóa phần tử:", updatedSlice)
}
