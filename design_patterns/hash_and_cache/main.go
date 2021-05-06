package main

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
)

// sometimes we generate things from one thing, we might re-generate the same thing.
// so we can simple create a cache for generated things.
//this is good to ensure the amount of data we are generating is manageable.

// return a hash value for an interface. the hash value is 16 bytes
func hash(obj interface{}) [16]byte {
	bytes, _ := json.Marshal(&obj)
	return md5.Sum(bytes)
}

func main() {
	// create a cache
	var cache = make(map[[16]byte]string)
	for i := 0; i < 10; i++ {
		// here we assume i is random number
		value := hash(i)
		cache[value] = string(rune(i))
	}
	fmt.Println(cache)
}
