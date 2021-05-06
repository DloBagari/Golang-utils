package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	t = t.Add(-6 * time.Hour)
	fmt.Println(t)
	//{"timestamp":"2021-04-15T19:21:32.076+03:00"}

}
