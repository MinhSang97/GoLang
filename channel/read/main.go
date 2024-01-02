package main

import "fmt"

func main() {
	c := make(chan int, 1)

	c <- 58

	d := <-c

	fmt.Println("number: ", d)
}
