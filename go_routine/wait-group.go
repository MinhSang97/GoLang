package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	// B1: tạo 1 biến có kiểu sync.WaitGroup: wg
	var wg sync.WaitGroup
	// B2: Wg.Add (số lượng goroutine cần phải đợi): wg.Add(4)
	wg.Add(4)
	// B3: ở mỗi goroutine, gọi methol wg.Done() trước khi return
	// B4: gọi mothod wg.Wait()
	go doSomething1(&wg)
	go doSomething2(&wg)
	go doSomething3(&wg)
	doSomething4(&wg)
	wg.Wait()
	log.Println(doSomething1)

}

func doSomething1(wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Println("số", i)
	}

	time.Sleep(time.Second * 2)
	fmt.Println("xong1")
	wg.Done()
}

func doSomething2(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("xong2")
	wg.Done()
}

func doSomething3(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("xong3")
	wg.Done()
}

func doSomething4(wg *sync.WaitGroup) {
	time.Sleep(time.Second * 2)
	fmt.Println("xong4")
	wg.Done()
}
