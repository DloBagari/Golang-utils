package main

import (
	"fmt"
	"math/rand"
	"time"
)

//wait for finished,  we wait for the goroutine to finished. of course in production we will use sync.WaitGroup for that
//here we use channel just to the the mechanic

func main() {
	// create a channel with signalling with out data, for that we use empty struct
	ch := make(chan struct{})
	go func() {
		time.Sleep(time.Duration(rand.Intn(6)) * time.Second)
		//close channel with will send a signal to receivers about the state change
		close(ch)
	}()
	//if withDate is true ==> the signal has data in it, which will be assigned to data
	// if withDate is false ==> data will be empty struct, because we have initialized ch with empty struct
	data, wd := <-ch
	fmt.Printf("%#v  %#v\n", data, wd)
}
