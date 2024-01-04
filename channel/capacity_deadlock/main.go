package main

import "fmt"

func main() {
	c := make(chan int)

	c <- 56

	fmt.Println("the end")
}
