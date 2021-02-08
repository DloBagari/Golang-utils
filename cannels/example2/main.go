package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

//wait for result from goroutine
func main() {
	runtime.GOMAXPROCS(1)
	p := make(chan string)
	go func() {
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		p <- "message"
	}()
	value := <-p
	fmt.Println(value)

}
