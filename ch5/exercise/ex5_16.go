package exercise

import "strings"

// Exercise 5.16: Write a variadic version of strings.Join.

// Variadic Join function using strings.Builder
func join(sep string, elems ...string) string {
	if len(elems) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.WriteString(elems[0])

	for _, elem := range elems[1:] {
		builder.WriteString(sep)
		builder.WriteString(elem)
	}

	return builder.String()
}
