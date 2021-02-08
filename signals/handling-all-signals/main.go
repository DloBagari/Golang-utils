package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig)
	go func() {
		for {
			rec := <-sig
			switch rec {
			case os.Interrupt:
				os.Exit(0)
			default:
				fmt.Println(rec)

			}
		}
	}()
	time.Sleep(10 * time.Second)
}
