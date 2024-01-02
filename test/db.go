package main

import "fmt"

func ping(ch chan string) {
  ch <- "ping"
} 

func pong(ch chan string) {
  msg := <- ch 
  fmt.Println(msg) // in ra "ping"
  ch <- "pong"
}

func main() {
  ch := make(chan string, 1)
  
  go ping(ch)
  go pong(ch)
  
  msg := <- ch 
  fmt.Println(msg) // in ra "pong"
}