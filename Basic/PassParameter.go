package main

import "fmt"

func main() {
	b := 10
	passByPointer(&b)
	println(b)
	slice1 := []int{1, 2, 3}
	passByValue2(&slice1)
	fmt.Println(slice1)

}

func passByValue(a int) {
	a = 100

}

func passByValue2(a *[]int) {
	*a = append(*a, 200)
}

func passByPointer(a *int) {
	*a = 100

}
