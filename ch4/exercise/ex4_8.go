package exercise

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

/*
Exercise 4.8: Modify charcount to count letters, digits, and so on in their Unicode categories,
using functions like unicode.IsLetter.
*/
func CharCount() {
	counts := make(map[rune]int)    // contagens de caracteres Unicode
	var utflen [utf8.UTFMax + 1]int // contagem dos comprimentos das codificações UTF-8
	invalid := 0                    // contagem de caracteres UTF-8 inválidos
	in := bufio.NewReader(os.Stdin)

	// Contadores para categorias
	lettersCount := 0
	digitsCount := 0
	othersCount := 0

	for {
		r, n, err := in.ReadRune() // retorna rune, nbytes, erro
		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}

		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}

		// Incrementa contagem por categoria
		if unicode.IsLetter(r) {
			lettersCount++
		} else if unicode.IsDigit(r) {
			digitsCount++
		} else {
			othersCount++
		}

		counts[r]++
		utflen[n]++
	}

	// Exibir resultados
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}

	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}

	fmt.Printf("\nTotal de letras: %d\n", lettersCount)
	fmt.Printf("Total de dígitos: %d\n", digitsCount)
	fmt.Printf("Total de outros caracteres: %d\n", othersCount)

	if invalid > 0 {
		fmt.Printf("\n%d caracteres UTF-8 inválidos\n", invalid)
	}
}
