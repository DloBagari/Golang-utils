package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

//when we work with multi threads and parallel we need to deal with synchronisation and orchestration for goroutines

func main() {
	// allow go to use 1 core, normally this should be called in an init() function of a package
	runtime.GOMAXPROCS(1)

	//create a waitGroup, the waitGroup can be used for anything which need to be tracked its termination. basically wait groups is used here for orchestration
	var wg sync.WaitGroup

	//set number of component that the wait group will track them
	wg.Add(2)

	//create two go routines
	start := time.Now()
	go func(value string) {
		count := 0
		for i := 0; i < 10000; i++ {
			fmt.Println(value)
			count += count
			//count down the wait group
		}
		wg.Done()
	}("go")

	go func(value string) {
		count := 0
		for i := 0; i < 10000; i++ {
			fmt.Println(value)
			count += count
			//count down the wait group
		}
		wg.Done()
	}("GO")
	wg.Wait()
	end := time.Since(start)
	fmt.Println(end)
}
