//////////////////////////////////////////////////////////////////////
//
// Your task is to change the code to limit the crawler to at most one
// page per second, while maintaining concurrency (in other words,
// Crawl() must be called concurrently)
//
// @hint: you can achieve this by adding 3 lines
//

package main

import (
	"fmt"
	"sync"
	"time"
)

// My First Solution - 1
// var mu sync.Mutex

// Second Solution - 1 
const rateLimit = 1 * time.Second
var throttle = time.Tick(rateLimit)

// Crawl uses `fetcher` from the `mockfetcher.go` file to imitate a
// real crawler. It crawls until the maximum depth has reached.
func Crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	if depth <= 0 {
		return
	}

	// My First Solution - 2
	// mu.Lock()
	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	mu.Unlock()
	// }()

	// Second Solution - 2
	<-throttle
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q\n", url, body)

	wg.Add(len(urls))
	for _, u := range urls {
		// Do not remove the `go` keyword, as Crawl() must be
		// called concurrently
		go Crawl(u, depth-1, wg)
	}
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	Crawl("http://golang.org/", 4, &wg)
	wg.Wait()
}
