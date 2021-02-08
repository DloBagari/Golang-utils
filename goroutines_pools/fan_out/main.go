package main

import (
	"fmt"
	"runtime"
	"time"
)

// we need to create 100 goroutines but we allow only 10 of them to run at the same time.
const (
	worker  = 10
	running = 100
)

func main() {
	runtime.GOMAXPROCS(2)
	workerChan := make(chan string, worker)
	runningChan := make(chan bool, running)

	for i := 0; i < worker; i++ {
		go func(workerId int) {
			runningChan <- true
			time.Sleep(time.Duration(1000) * time.Millisecond)
			workerChan <- fmt.Sprintf("message from worker: %d", workerId+1)
			fmt.Printf("worker number %d has done its work\n", workerId+1)
			<-runningChan
		}(i)
	}

	for i := 0; i < worker; i++ {
		fmt.Println(<-workerChan)
	}
}
