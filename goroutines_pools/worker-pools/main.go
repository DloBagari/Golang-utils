package main

import (
	"fmt"
	"sync"
)

type Client struct {
	Id     int
	Number int
}

type Data struct {
	// if we make this as pointer to client, then all the Data passed to data channels will have the same client
	Job    Client
	Result int
}

var (
	size    = 100
	clients = make(chan Client, size)
	data    = make(chan Data, size)
)

func Workers(w *sync.WaitGroup) {
	for c := range clients {
		// NOW: c is in the for literal scoop, when we pass c as pointer to any channel, the value of the c will be the value of what c is in the for loop hen we read c from the channel.
		r := c.Number * c.Number
		data <- Data{
			Job:    c,
			Result: r,
		}
	}
	w.Done()
}

func CreateJobs(n int) {
	for i := 0; i < n; i++ {
		c := Client{
			Id:     i,
			Number: i,
		}
		clients <- c
	}
	close(clients)
}

func CreateWorkers(n int) {
	var wg sync.WaitGroup
	for i := 0; i < n; i++ {
		wg.Add(1)
		go Workers(&wg)
	}
	wg.Wait()
	close(data)
}

func main() {
	go CreateJobs(10000)
	finished := make(chan interface{})
	go func() {
		for d := range data {
			fmt.Printf("ClientId: %d\t int: %d\t result: %d\t\n", d.Job.Id, d.Job.Number, d.Result)
		}
		finished <- true
	}()
	CreateWorkers(200)
	<-finished
}
