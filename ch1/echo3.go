package ch1

import (
	"os"
	"strings"
)

func Echo3() string {
	return strings.Join(os.Args[1:], " - ")
}
