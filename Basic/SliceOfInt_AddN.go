// package main

// import "fmt"

// func main() {
// 	// Tạo một slice ban đầu
// 	slice := []int{1, 2, 3, 4, 5}
// 	fmt.Println("Slice ban đầu:", slice)
// 	fmt.Println("-------------")

// 	// Thêm một phần tử vào vị trí n (vd: vị trí thứ 2)
// 	n := 2
// 	newElement := 10

// 	// Tách slice thành hai phần
// 	firstPart := slice[:n]
// 	secondPart := slice[n:]

// 	// Tạo slice mới kết hợp cả phần đầu, phần tử mới và phần cuối
// 	updatedSlice := append(firstPart, newElement)
// 	updatedSlice = append(updatedSlice, secondPart...)

// 	fmt.Println("Slice sau khi thêm phần tử:", updatedSlice)
// }

package main

import "fmt"

func main() {

	// Tạo một slice ban đầu
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println("Slice ban đầu:", slice)
	fmt.Println("-------------")

	// Nhập giá trị n từ người dùng
	var n int
	fmt.Print("Nhập giá trị n: ")
	fmt.Scanln(&n)

	// Kiểm tra xem n có hợp lệ không
	if n < 0 || n >= len(slice) {
		fmt.Println("n không hợp lệ.")
		return
	}

	// Nhập giá trị n từ người dùng
	var newElement int
	fmt.Print("Nhập giá trị bạn muốn thêm vào Slice: ")
	fmt.Scanln(&newElement)

	// Tạo slice mới
	updatedSlice := make([]int, len(slice)+1)

	// Copy các phần tử từ slice ban đầu đến phần đầu của updatedSlice
	copy(updatedSlice[:n], slice[:n])

	// Copy phần tử mới vào updatedSlice
	updatedSlice[n] = newElement

	// Copy các phần tử từ slice ban đầu sau vị trí n đến updatedSlice
	copy(updatedSlice[n+1:], slice[n:])

	fmt.Println("Slice sau khi thêm phần tử:", updatedSlice)
}
