package main

import (
	"fmt"
	"github.com/Golang-utils/async_client_requests"
	"net/http"
)

func main() {
	urls := []string{
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"github.com",
		"google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"https://www.google.com",
		"https://www.golang.org",
		"https://www.github.com",
		"google.com",
		"https://www.golang.org",
		"https://www.github.com",
	}
	c := async_client_requests.NewClient(http.DefaultClient, len(urls))
	async_client_requests.FetchAll(urls, c)
	//Go’s select lets you wait on multiple channel operations
	// since we will have N request, So we wait for N response or error
	//by this way we will be waiting for all goroutines to be finished
	for i := 0; i < len(urls); i++ {
		select {
		case resp := <-c.Resp:
			fmt.Printf("status received from %s: %d\n", resp.Request.URL, resp.StatusCode)
		case err := <-c.Err:
			fmt.Printf("error received %s\n", err)
		}
	}
}
