package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	data := make(map[string]string)
	var mx sync.Mutex
	// create conditional variable:
	c := sync.NewCond(&mx)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// lock on the condition
			c.L.Lock()
			//if there is no items in data wait
			if len(data) == 0 {
				// wait until it receives a signal
				c.Wait()
			}
			fmt.Println(data["key"])
			c.L.Unlock()
		}()
	}
	//create goroutine which will add item to data and send signal
	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(4 * time.Second)
		c.L.Lock()
		data["key"] = "value"
		c.Broadcast()
		c.L.Unlock()
	}()
	wg.Wait()
}
