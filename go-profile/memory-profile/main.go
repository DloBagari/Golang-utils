package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"runtime/pprof"
	"time"
)

func n1(n int) int {
	fn := make(map[int]int)
	for i := 0; i < n; i++ {
		var f int
		if i <= 2 {
			fn[i] = 1
			continue
		}
		f = fn[i-1] + fn[i-2]
		d := math.Floor(float64(17*n/2 + 1))
		math.Floor(d*d/2 + 1*d)
		fn[i] = f
	}
	time.Sleep(100 * time.Millisecond)
	return fn[n]
}

func n2(n int) bool {
	k := math.Floor(float64(n/2 + 1))
	for i := 2; i < int(k); i++ {
		if (n % i) == 0 {
			return false
		}
	}
	return true
}

func n3(n int) {
	for i := 2; i < n; i++ {
		if n2(i) {
		}
	}

}

func main() {
	// create file to hold profile data for cpu
	file, _ := os.Create("/home/dlo/test/memory_profile")
	defer file.Close()
	defer pprof.StopCPUProfile()
	fmt.Println(n1(1000000))
	n3(100000)
	if err := pprof.WriteHeapProfile(file); err != nil {
		log.Fatal(err)
	}
}
