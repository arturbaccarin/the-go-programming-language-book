package exercise

/*
Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer instead of string
concatenation.
*/
import (
	"bytes"
	"strings"
)

// comma formats a number with commas as thousands separators.
func commaEnhanced(s string) string {
	var integerPart, decimalPart string

	if strings.Contains(s, ".") {
		integerPart = strings.Split(s, ".")[0]
		decimalPart = strings.Split(s, ".")[1]
	} else {
		integerPart = s
	}

	var buf bytes.Buffer
	var signal string

	n := len(integerPart)

	if integerPart[0] == '-' {
		signal = string(s[0])
		n--
		integerPart = integerPart[1:]
	}

	buf.WriteString(signal)

	// Iterar sobre o string
	for i, c := range integerPart {
		// Adiciona uma vírgula antes de adicionar um novo dígito se estiver na posição certa
		if i > 0 && (n-i)%3 == 0 {
			buf.WriteByte(',')
		}
		// Adiciona o caractere atual ao buffer
		buf.WriteByte(byte(c))
	}

	if decimalPart != "" {
		buf.WriteString("." + decimalPart)
	}

	return buf.String()
}
