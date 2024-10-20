//////////////////////////////////////////////////////////////////////
//
// Given is a producer-consumer scenario, where a producer reads in
// tweets from a mockstream and a consumer is processing the
// data. Your task is to change the code so that the producer as well
// as the consumer can run concurrently
//

package main

import (
	"fmt"
	"time"
)

func producer(stream Stream, ch chan *Tweet, quit chan bool) {
	defer close(quit)
	for {
		tweet, err := stream.Next()
		if err == ErrEOF {
			return
		}

		ch <- tweet
	}
}

func consumer(tweets []*Tweet) {
	for _, t := range tweets {
		if t.IsTalkingAboutGo() {
			fmt.Println(t.Username, "\ttweets about golang")
		} else {
			fmt.Println(t.Username, "\tdoes not tweet about golang")
		}
	}
}

func main() {
	start := time.Now()
	stream := GetMockStream()

	ch := make(chan *Tweet, 10)
	quit := make(chan bool)

	// Producer
	go producer(stream, ch, quit)

	// Consumer
	loop:
		for {
			select {
			case t := <-ch:
				consumer([]*Tweet{t})
			case <-quit:
				break loop
			}
		}

	fmt.Printf("Process took %s\n", time.Since(start))
}
