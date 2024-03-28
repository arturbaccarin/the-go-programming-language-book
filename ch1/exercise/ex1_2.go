package exercise

import (
	"fmt"
	"os"
)

// Exercise 1.2 - Modify the echo program to print the index and value of each of its arguments, one per line.
func Exercise1_2() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
