package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		i := 0
		for {
			fmt.Println("number", i)
			i++
			time.Sleep(time.Millisecond * 10)
		}
	}()

	// Để giữ chương trình chính không kết thúc ngay lập tức
	select {}
}
