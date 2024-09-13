package exercise

/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string
concatenation.
*/
import (
	"bytes"
	"fmt"
)

// comma formats a number with commas as thousands separators.
func comma(s string) string {
	var buf bytes.Buffer
	n := len(s)

	// Iterar sobre o string
	for i, c := range s {
		// Adiciona uma vírgula antes de adicionar um novo dígito se estiver na posição certa
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		// Adiciona o caractere atual ao buffer
		buf.WriteByte(byte(c))
	}

	return buf.String()
}

func main() {
	fmt.Println(comma("1234567890")) // Output: "1,234,567,890"
	fmt.Println(comma("1234"))       // Output: "1,234"
	fmt.Println(comma("12"))         // Output: "12"
}
