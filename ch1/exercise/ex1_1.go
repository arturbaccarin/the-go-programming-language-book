package exercise

import (
	"fmt"
	"os"
	"strings"
)

func Exercise1_1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
