package main

import (
	"fmt"
	"github.com/Golang-utils/sync_once"
)

func main() {
	// create once instance of logrus
	if err := sync_once.UseLog(); err != nil {
		fmt.Println(err)
	}

	//try to create another instance
	if err := sync_once.UseLog(); err != nil {
		fmt.Println(err)
	}
}
