package main

import (
	"fmt"
	"time"
)

// checkMultipleReceivers
// o/p: only one Hi, Yashu will be printed
// Learn: one channel message will be read by only one active listener 

func listen(ch chan string) {
	v := <-ch
	fmt.Println(v)
}

func checkMultipleReceivers() {
	ch := make(chan string)
	go listen(ch)
	go listen(ch)

	ch <- "Hi, Yashu"
	time.Sleep(10 * time.Second)
}


// unbufferedChannels 
// sending in mulitple messages on to a unbuffered channel, without readers on concurrent
// deadlock senerio
func unbufferedChannels() {
	ch := make(chan string)
	ch <- "Hi, Yashu"
	ch <- "Hi, Yashu"
	time.Sleep(10 * time.Second)	
}

func main() {
	// 1.
	// checkMultipleReceivers()

	// 2.
	unbufferedChannels()
}