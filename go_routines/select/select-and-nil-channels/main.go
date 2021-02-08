package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//to block or disable on branch in select, we set the channel to nil
// VERY IMPORTANT, if wg has two goroutines and both go routines are locked with in their select statement ==> the program will panic
//this because all goroutines the become asleep.

func add(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	sum := 0
	// timer once it is called t.C it  will count until the given time and return a time chan.
	t := time.NewTimer(1 * time.Second)
	for {
		select {
		case i := <-c:
			sum += i
		case <-t.C:
			c = nil
			fmt.Println(sum)
			// if we remove this ==> this select statement will be locked, since all its branches are locked.
			t.Reset(time.Second)
		}
	}
}

func send(c chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		c <- rand.Intn(1000000000000)
	}
}

func main() {
	var wg sync.WaitGroup
	c := make(chan int)
	wg.Add(2)
	go add(c, &wg)
	go send(c, &wg)
	wg.Wait()
	time.Sleep(5 * time.Second)

}
