package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int, 10)

	go writeToChannel(c)

	// khi đọc ở 1 empty channel thì nó sẽ làm cái current go routine sleep
	// var d int

	// for i := 0; i < 11; i++ {
	// 	d = <-c
	// 	fmt.Println("number: ", d)
	// }

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

}
