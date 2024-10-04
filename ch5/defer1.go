package ch5

import "fmt"

func main() {
	f(3)
}

// When a panic occurs, all defer red functions are run in reverse order, starting wit h thos e of the
// top most function on the stack and pro ceeding up to main
func f(x int) {
	fmt.Printf("f(%d)\n", x+0/x) // panics if x == 0
	defer fmt.Printf("defer %d\n", x)
	f(x - 1)
}

/*
f(3)
f(2)
f(1)
defer 1
defer 2
defer 3
*/
