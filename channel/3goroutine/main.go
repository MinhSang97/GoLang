package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool, 3)
	numGoroutines := 3

	go doSomething1(done)

	go doSomething2(done)

	go func() {
		doSomething3()
		done <- true
	}()

	for i := 0; i < numGoroutines; i++ {
		<-done
	}

	close(done)

	// for j := range done {
	// 	fmt.Println("Done", j)
	// }

	fmt.Println("All goroutines are done.")
}

func doSomething1(ch chan bool) {
	defer func() {
		ch <- true
	}()
	for i := 0; i < 10; i++ {
		fmt.Println("số", i)
	}
	time.Sleep(time.Second * 2)
	fmt.Println("xong1")

}

func doSomething2(ch chan bool) {
	time.Sleep(time.Second * 2)
	fmt.Println("xong2")
	ch <- true
}

func doSomething3() {
	time.Sleep(time.Second * 2)
	fmt.Println("xong3")
}
