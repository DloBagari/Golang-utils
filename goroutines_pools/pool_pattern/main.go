package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//here we create a pool of goroutine which will have 2 goroutines
//they will wait to get tasks and process them

const (
	workers = 2
	work    = 10
)

func main() {
	ch := make(chan int)
	runtime.GOMAXPROCS(workers)
	for i := 0; i < workers; i++ {
		go func(workerId int) {
			for w := range ch {
				fmt.Printf("worker %d took a task: %d\n", workerId+1, w)
				time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
			}
		}(i)
	}
	for i := 0; i < work; i++ {
		ch <- i
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
	}
}
