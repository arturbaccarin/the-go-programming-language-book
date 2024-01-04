package ch1

import (
	"os"
)

func Echo2() string {
	s, sep := "", ""

	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}

	return s
}
