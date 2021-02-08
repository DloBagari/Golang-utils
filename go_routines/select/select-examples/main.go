package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// all the cases in select statement are not evaluated sequentially, they are evaluated simultaneously
// if none of the select statement are ready, the select statement will block until one of its statement become ready
// if multiple channels in the select statement are ready, then the go runtime will make a random selection.
// select statement with time.After() can be used to timeout a goroutine.

// the zero value of the channel type is nil. var c chan string : here c is nil. if we try to close a nil channel the program will panic
// if we send a message to a closed channel ==> program will panic
//if we read from a closed channel we will get  the zero value of the type of that channel.
// so after closing a channel we can not longer write to it, but we still can read from it.
// in order to be able to close a channel, the channel must not be receive-only. for example if we pass to a function a channel  <-chan, and we try to close this change
//we will fail

// if we read/write to nil channel, the channel will be blocked: this can be useful if we want to disable a branch in
//the select statement by assigning nil value to the channel variable

func gen(min, max int, createNumber chan<- int, end <-chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case createNumber <- rand.Intn(max-min) + min:
		case <-end:
			fmt.Println("returned")
			return
		case <-time.After(4 * time.Second):
			fmt.Println("rec time")
		}
	}
}

func main() {
	var wg sync.WaitGroup
	createNumber := make(chan int)
	end := make(chan bool, 1)
	wg.Add(1)
	// wg most be passed as pointer, if we pass as value ==> we will get deadlock, you know why?? :)
	go gen(20, 120, createNumber, end, &wg)
	for i := 0; i < 1000; i++ {
		fmt.Println(<-createNumber)
	}
	time.Sleep(5 * time.Second)
	end <- true
	wg.Wait()
	close(createNumber)
	close(end)
}
