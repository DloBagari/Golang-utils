package main

import (
	"fmt"
	"runtime"
	"time"
)

// closure variable, is a variable which is used in a goroutine, where that variable is not passed to the goroutine.
//when a goroutine is referencing a variable outside its body, this variable is evaluated at the time when the goroutine
//is executed.

func main() {
	//here all goroutines will be waiting the go scheduler to execute them, once the scheduler execute them, the for loop might be
	//finished and the value of if will be 201, so when the goroutine evaluate the i variable, multiple go routine might get the save value of i
	for i := 0; i <= 200; i++ {
		go func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println("\n-----------------")
	//the solution for this is to intermediate variable which will be used inside the go routine
	for i := 0; i <= 200; i++ {
		i := i
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)

	//another issue, notice here the condition in for loop is always true. since v is byte, adding 1 to 255 the byte i will overflow and become 0
	var i byte
	go func() {
		for i = 0; i <= 255; i++ {
		}
	}()
	fmt.Println("calling go scheduler to execute another goroutine")
	runtime.Gosched()
	fmt.Println("calling go garbage collection")
	runtime.GC()
	fmt.Println("done")

}
