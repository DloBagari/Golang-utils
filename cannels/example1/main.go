package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

// wait for task to give to the goroutine
func main() {
	p := make(chan string)
	runtime.GOMAXPROCS(1)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		value := <-p
		fmt.Println("value from goroutine: ", value)
		wg.Done()
	}()
	time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
	p <- "message"
	wg.Wait()
}
