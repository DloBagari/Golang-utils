package main

import (
	"fmt"
	"sync"
	"time"
)

// IMPORTANT if we use 'var read chan int' the channel will be nil, and if we read/write to it we will get panic
var read = make(chan int)
var write = make(chan int)

func readValue() int {
	return <-read
}

func writeValue(i int) {
	write <- i
}

func monitor() {
	var value int
	for {
		select {
		case v := <-write:
			value = v
		// how many time this will be executed? well, it will executed only when there is a read form the read changed
		//because in channels, the receiving the happening before sending. so if there is not receiving, than there will not be sending
		case read <- value:
			fmt.Println("added")
		}
	}
}

func main() {
	go monitor()
	time.Sleep(time.Second)
	var wgg sync.WaitGroup
	for i := 0; i < 3; i++ {
		wgg.Add(1)
		go func() {
			defer wgg.Done()
			writeValue(i)
		}()

	}
	wgg.Wait()
	fmt.Println(readValue())
}
