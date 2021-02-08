package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

//pressing ctrl-c creates a SIGINT signal. in go this is os.Interrupt
// signals are the way to handle asynchronous events on a unix system.

func main() {
	//create a signal channel with buffer 1
	sigs := make(chan os.Signal, 1)
	//define the signals we are interested in
	signal.Notify(sigs, os.Interrupt, syscall.SIGXCPU)
	go func() {
		for {
			rec := <-sigs
			switch rec {
			case os.Interrupt:
				fmt.Println(rec)
				os.Exit(0)
			case syscall.SIGXCPU:
				fmt.Println(rec)

			}
		}
	}()
	for {
		fmt.Println(".")
		time.Sleep(10 * time.Second)
	}

}
