package main

import (
	"fmt"
	"github.com/pkg/profile"
	"math"
	"time"
)

// go tool pprof ....
//go tool pprof --http:localhost:8080 'path to data'

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
	defer profile.Start(profile.ProfilePath("/home/dlo/test/cpu2_profile")).Stop()
	// for memory: defer profile.Start(profile.MemProfile).Stop()
	// result will be stored in /tmp
	fmt.Println(n1(1000000))
	n3(100000)

}
