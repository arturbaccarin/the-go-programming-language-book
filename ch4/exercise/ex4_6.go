package exercise

import (
	"bytes"
	"unicode"
)

/*
Exercise 4.6: Write an in-place function that squashes each run of adjacent Unicode spaces
(see unicode.IsSpace) in a UTF-8-encoded []byte slice into a single ASCII space.
*/
func Squash(input []byte) []byte {
	var result bytes.Buffer
	prevSpace := false

	for _, b := range input {
		if unicode.IsSpace(rune(b)) {
			if !prevSpace {
				result.WriteByte(' ') // Write only one space for consecutive spaces
				prevSpace = true
			}
		} else {
			result.WriteByte(b) // Write the non-space byte
			prevSpace = false
		}
	}

	return result.Bytes()
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
