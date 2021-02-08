package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// read coordinates from file and plot them with Glot.
//Glot is external modules it needs Gnuplot packages for linux
// sudo apt-get install -y gnuplot
// go get github.com/Arafatk/glot

func main() {
	csvFilePath := ""
	//find if the file is exist using os.Stat
	_, err := os.Stat(csvFilePath)
	if err != nil {
		log.Fatal(err)
	}
	//open the file
	f, err := os.Open(csvFilePath)
	defer f.Close()
	reader := csv.NewReader(f)
	reader.FieldsPerRecord = -1
	//allRecords is [][]string
	allRecords, err := reader.ReadAll()
	fmt.Println(allRecords)
}
