package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)

	go writeToChannel(c)

	for j := range c {
		fmt.Println("number: ", j)

	}

}

func writeToChannel(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		fmt.Println("số", i)
	}
	time.Sleep(time.Millisecond * 200)
	fmt.Println("truyen xong")
	close(ch)

}
