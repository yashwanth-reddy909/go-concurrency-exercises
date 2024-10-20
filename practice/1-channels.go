package main

import (
	"fmt"
	"time"
)

// o/p: only one Hi, Yashu will be printed
// Learn: one channel message will be read by only one active listener 

func listen(ch chan string) {
	v := <-ch
	fmt.Println(v)
}

func main() {
	ch := make(chan string)
	go listen(ch)
	go listen(ch)

	ch <- "Hi, Yashu"
	time.Sleep(10 * time.Second)
}
