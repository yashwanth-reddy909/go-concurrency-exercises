//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  atomic.Int64 // in seconds
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	if u.IsPremium {
		process()
		return true
	}

	done := make(chan bool)

	// process runs on seperate go-routine
	go func() {
		process()
		fmt.Println("process done", u.ID)
		done <- true
	}()

	// ticker runs on seperate go-routine
	// this will update TimeUsed on every Tick of 1 second
	ticker := time.NewTicker(time.Second * 1)
	go func() {
		for range ticker.C {
			u.TimeUsed.Add(1)
			if u.TimeUsed.Load() >= 10 {
				done <- false
			}
		}
	}()

	select {
	case res := <-done:
		ticker.Stop()
		return res
	}

	// Solution 1:
	// if (u.IsPremium) {
	// 	process()
	// 	return true
	// }

	// timeLimit := make(chan bool)
	// done := make(chan bool)
	// go func() {
	// 	time.Sleep(10 * time.Second)
	// 	timeLimit <- true
	// }()

	// go func() {
	// 	process()
	// 	done <- true
	// }()

	// select {
	// case <-timeLimit:
	// 	return false
	// case <-done:
	// 	return true
	// }
}

func main() {
	RunMockServer()
}
