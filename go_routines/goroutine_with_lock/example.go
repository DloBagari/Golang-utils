package main

import (
	"fmt"
	"sync"
)

// note: if we have a global variable and it not used go will not complain
var (
	//normally we add mutex in a struct and ensure there is only one instance of it is exist. singleton
	//sync.Mutex will block read and writes and allow only one goroutine to do read and write
	//but the read operation does not change anything in memory, so if we want to only lock the write operation we can use
	// sync.RWMutex and when we modify a value then we can lock the read with  mutex.RLock.try that
	//basically any goroutine can read when no write operation is taking place. this is good for slices and maps
	//we use RWMutex only if we have multiple goroutines reading the same object.
	//BUT remember this cause latency :(

	//go has race condition detector to find out if two goroutines are trying to race about a memory location
	// to appy race detection on your code on build time, pass -race parameter to build command
	//also we should run te race detector on testing :)
	//for map access the race detection is automatically applied to the map on run time. if teo goroutines try to
	// modify the same map at the same time==> at run time go will detect race condition:)
	// if we use MAXPROCS=1 ./executable_go_app we might not detect race because this will not run in parallel.
	//in docker env to run a go app with single core add MAXPROCS=1 before the app exec command example: MAXPROCS=2 ./app
	mutex sync.Mutex
	count int
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {

		for i := 0; i < 2; i++ {
			mutex.Lock()
			{
				value := count
				fmt.Println(value)
				value++
				count = value
			}
			mutex.Unlock()

		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 2; i++ {
			mutex.Lock()
			{
				value := count
				fmt.Println(value)
				// ++ is read  modify and write
				//fmt.printline will generate syscall which will let the goroutine status to be wait.
				// this will let the other gorotinue to overwrite count, to avoid that we use mutex to lock and allow only one gorotine to execute a set of instrctures
				value++
				count = value
			}
			mutex.Unlock()
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("Count: ", count)

}
