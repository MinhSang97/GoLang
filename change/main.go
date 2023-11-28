package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6 ,2,2,7,0,34,234,252,16,26,167,15}
	b := []int{3, 5, 7, 9, 11}

	fmt.Println("Slice a", a)
	fmt.Println("===================")
	fmt.Println("Slice b", b)
	fmt.Println("===================")

	for _, elementA := range a {
		for _, elementB := range b {
			if elementA == elementB {
				fmt.Printf("%d thuộc cả a và b\n", elementA)
				break
			} else {
				fmt.Printf("%d không thuộc cả a và b\n", elementA)
				break
			}
		}
	}
}
