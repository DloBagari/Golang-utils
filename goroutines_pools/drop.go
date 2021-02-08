package main

import (
	"fmt"
	"math/rand"
	"time"
)

// how to drop adding sending data when the channel buffer is full?
const (
	worker = 100
	caps   = 5
)

func main() {
	ch := make(chan string, caps)
	go func() {
		for p := range ch {
			//fmt.Printf("got message %s\n", p)
			_ = p
			time.Sleep(time.Duration(rand.Intn(10)) * time.Millisecond)

		}
	}()

	for i := 0; i < worker; i++ {
		//The select statement lets a goroutine wait on multiple communication operations.
		//if multiple channels are ready, it will choose one randomly
		time.Sleep(time.Duration(rand.Intn(5)) * time.Millisecond)
		select {
		//try to send data to channel, if the buffer is full, this case will fail and the default one will be executed
		case ch <- fmt.Sprintf("mmessage: %d", i):
			fmt.Printf("message %d is sent\n", i)
		default:
			fmt.Printf("message %d is dropped\n", i)
		}
	}
}
