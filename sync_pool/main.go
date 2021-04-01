package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sync"
	"time"
)

//sync.pool can be used to create expensive resources, where those resource can be pull from a pool, and once they are used they can be put back to the pool
//for example: create a pool of byte buffer.
var bufferPool = sync.Pool{
	//this method will be called when there is no free buffer in the pool
	// the value for the new is a function, which returns a byte buffer
	New: func() interface{} {
		fmt.Println("creating a new buffer")
		return new(bytes.Buffer)
	},
}

func log(w io.Writer, message string) {
	buff := bufferPool.Get().(*bytes.Buffer)
	//reset the buffer
	buff.Reset()
	//there to the buffer
	buff.WriteString(time.Now().Format("15:04:05"))
	buff.WriteString(":")
	buff.WriteString(message)
	buff.WriteString("\n")
	w.Write(buff.Bytes())
	//put the buffer back to the pool
	bufferPool.Put(buff)
}

func main() {
	log(os.Stdout, "fist message")
	log(os.Stdout, "second message")
	log(os.Stdout, "third message")
}
