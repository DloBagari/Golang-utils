package main

import (
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/trace"
)

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	//print in MB
	fmt.Println("mem.Alloc: ", mem.Alloc/(1024*1024))
	fmt.Println("mem.TotalAlloc: ", mem.TotalAlloc/(1024*1024))
	fmt.Println("mem.HeapAlloc: ", mem.HeapAlloc/(1024*1024))
	fmt.Println("mem.MunGC: ", mem.NumGC)
	fmt.Println("-----------------------")
}

// google-chrome --enable-blink-features=ShadowDOMV0,CustomElementsV0,HTMLImports
// go tool trace home/dlo/test/traceFile.out

func main() {
	f, _ := os.Create("/home/dlo/test/traceFile.out")
	defer f.Close()
	// start the trace: to collect data
	if err := trace.Start(f); err != nil {
		log.Fatal(err)
	}
	defer trace.Stop()

	var mem runtime.MemStats
	printStats(mem)
	for i := 0; i < 3; i++ {
		// allocate some memory
		s := make([]byte, 50000000)
		_ = s
	}
	printStats(mem)

	for i := 0; i < 20; i++ {
		// allocate some memory
		s := make([]byte, 50000000)
		_ = s
	}
	printStats(mem)

}
