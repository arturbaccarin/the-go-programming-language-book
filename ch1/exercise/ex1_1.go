package exercise

import (
	"fmt"
	"os"
	"strings"
)

// Exercise 1.1 - Modify the echo program to also print os.Args[0], the name of the command that invoked it.
func Exercise1_1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
