package main

import (
	"fmt"
	"sync"
)

// pipelines has stages, where each stage is represented by a goroutine (we can scale the number of goroutines in each stage).
// example: create the following pipeline: generate number ---> square number ---> print number

func generator(r int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		for i := 0; i < r; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func square(numbers <-chan int, wg *sync.WaitGroup) <-chan int {
	out := make(chan int)
	go func() {
		defer wg.Done()
		for i := range numbers {
			out <- i * i
		}
		close(out)
	}()
	return out
}

func displayNumber(numbers <-chan int, wg *sync.WaitGroup) {
	go func() {
		defer wg.Done()
		for i := range numbers {
			fmt.Println(i)
		}
	}()

}

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	displayNumber(square(generator(10, &wg), &wg), &wg)
	wg.Wait()
}
