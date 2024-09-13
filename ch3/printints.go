package ch3

import (
	"bytes"
	"fmt"
)

/*
The bytes package provides  the Buffer type  for  efﬁcient manipu lat ion of byte slices.  A
Buffer st arts out emp ty but grows as dat a of typ es li ke string, byte, and []byte are writt en
to it.  As the example below shows, a bytes.Buffer var iable requires no initializat ion because
its zero value is usable:
*/

// intsToString is like fmt.Sprintf(values) but adds commas.
func intsToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}

	buf.WriteByte(']')
	return buf.String()
}

func main() {
	fmt.Println(intsToString([]int{1, 2, 3})) // "[1, 2, 3]"
}
