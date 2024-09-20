package exercise

/*
Exercise 4.7: Modify reverse to reverse the characters of a []byte slice that represents a UTF-8-encoded string, in place. Can you do it without allocating new memory?

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
*/
package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"
)

func reverseUTF8(b []byte) {
	// Step 1: Reverse runes
	for i := 0; i < len(b)/2; i++ {
		b[i], b[len(b)-i-1] = b[len(b)-i-1], b[i]
	}

	// Step 2: Reassemble UTF-8 encoded string
	var result bytes.Buffer
	for len(b) > 0 {
		r, size := utf8.DecodeLastRune(b[:len(b)-utf8.UTFMax])
		result.WriteRune(r)
		b = b[:len(b)-size]
	}
}

func main() {
	input := []byte("Hello, 世界!")
	fmt.Printf("Before reverse:\t%s\n", input)
	reverseUTF8(input)
	fmt.Printf("After reverse:\t%s\n", input)
}
