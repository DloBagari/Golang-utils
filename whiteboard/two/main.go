package main

import (
	crand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"

	length := 10
	for i := 2; i < length; i++ {
		v, _ := crand.Int(crand.Reader, big.NewInt(int64(len(digits))))
		fmt.Println(v)
	}
}
