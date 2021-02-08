package main

import (
	"io/ioutil"
	"log"
	"os"
)

func main() {
	filePath := ""
	err := ioutil.WriteFile(filePath, []byte("test to write"), os.ModePerm) // or 0644
	if err != nil {
		log.Fatal(err)
	}
}
