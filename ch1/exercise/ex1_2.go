package exercise

import (
	"fmt"
	"os"
)

func Exercise1_2() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}
