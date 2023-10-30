// package main

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"sync"
// 	"time"
// )

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	var wg sync.WaitGroup
// 	wg.Add(4)

// 	go func(ctx context.Context) {
// 		defer wg.Done()

// 		for i := 0; i < 10; i++ {
// 			fmt.Println("số", i)
// 		}

// 		select {
// 		case <-time.After(time.Second * 2):
// 			fmt.Println("xong1")
// 		case <-ctx.Done():
// 			return
// 		}
// 	}(ctx)

// 	go func(ctx context.Context) {
// 		select {
// 		case <-time.After(time.Second * 210):
// 			fmt.Println("xong2")
// 		case <-ctx.Done():
// 			return
// 		}
// 		wg.Done()
// 	}(ctx)

// 	go func(ctx context.Context) {
// 		select {
// 		case <-time.After(time.Second * 360):
// 			fmt.Println("xong3")
// 		case <-ctx.Done():
// 			return
// 		}
// 		wg.Done()
// 	}(ctx)

// 	go func(ctx context.Context) {
// 		select {
// 		case <-time.After(time.Second * 2):
// 			fmt.Println("xong4")
// 		case <-ctx.Done():
// 			return
// 		}
// 		wg.Done()
// 	}(ctx)

// 	wg.Wait()
// 	log.Println(`Hoàn thành lúc`, time.Now())
// }
