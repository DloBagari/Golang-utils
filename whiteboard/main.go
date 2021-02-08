package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	for i := 0; i < 100000; i++ {
		if i/5 == 0 {
			s := i
			_ = s
		} else {
			s := i
			_ = s
		}
	}
	fmt.Println(time.Now().Sub(start))
}
