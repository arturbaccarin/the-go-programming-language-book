package ch5

import (
	"os"
	"runtime"
)

// For diagnostic purposes, the runtime package lets the programmer dump the stack using the
// same machinery. By deferring a call to printStack in main,

func main() {
	defer printStack()
	f(3)
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}
