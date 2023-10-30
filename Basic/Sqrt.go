package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Print("Type number a: ")
	fmt.Scanln(&a)
	fmt.Print("Type number b: ")
	fmt.Scanln(&b)
	fmt.Println("Your numbers are:", a, "and", b)
	
	// Gọi hàm results để tính và in ra kết quả
	results(float64(a), float64(b))
}

func results(a, b float64) {
	// Tính căn bậc hai của a^2 + b^2
	sumOfSquares := math.Pow(a, 2) + math.Pow(b, 2)
	result := math.Sqrt(sumOfSquares)
	fmt.Println("Square root of the sum of squares:", result)
}
