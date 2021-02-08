package main

import (
	"context"
	"flag"
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

func main() {
	duration := flag.Duration("timeout", 500*time.Millisecond, "timeout in milliseconds")
	ctx, cancel := context.WithTimeout(context.Background(), *duration)
	g, ctx := errgroup.WithContext(ctx)
	var urls = []string{
		"http://www.golang.org/",
		"https://enterprise.fyde.com/api/v1/users",
		"http://www.google.com/",
		"http://www.hotmail.com/",
		"http://www.facebook.com/",
		"http://www.github.com/",
	}

	for i, url := range urls {
		url := url
		i := i
		g.Go(func() error {
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
				res, err := http.Get(url)
				fmt.Printf("%+v\n\n\n------------%d\n", res.Header, i)
				if err == nil {
					res.Body.Close()
				}
				if res.StatusCode == 401 {
					cancel()
					return ctx.Err()
				}
			}
			return nil
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("Successfully fetched all URLs.")
	} else {
		fmt.Println(222)
	}

}
