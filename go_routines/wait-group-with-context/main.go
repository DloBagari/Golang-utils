package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), 4*time.Second)
	defer cancel()
	var urls = []string{
		"http://www.golang.org/",
		"https://enterprise.fyde.com/api/v1/users",
		"http://www.google.com/",
		"http://www.hotmail.com/",
		"http://www.facebook.com/",
		"http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
		"http://www.github.com/", "http://www.github.com/", "http://www.github.com/", "http://www.github.com/",
	}
	a := false
	for i, _ := range urls {
		if a {
			break
		}
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return // Error somewhere, terminate
			default:

			}
			fmt.Println(i)
			if i == 2 {
				fmt.Println(i)
				a = true
				cancel()
				return
			}

		}(i)
	}
	wg.Wait()

}
