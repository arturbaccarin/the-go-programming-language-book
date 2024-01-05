package exercise

import (
	"bufio"
	"fmt"
	"os"
)

// Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated line occurs.
func Exercise1_4() {
	counts := make(map[string]map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts, "stdin")
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, arg)
			f.Close()
		}
	}

	for fileName, lines := range counts {
		for line, n := range lines {
			if n > 1 {
				fmt.Println(fileName)
				fmt.Printf("%d\t%s\n", n, line)
			}
		}
	}
}

func countLines(f *os.File, counts map[string]map[string]int, fileName string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		_, ok := counts[fileName]
		if !ok {
			counts[fileName] = map[string]int{input.Text(): 1}
			continue
		}

		counts[fileName][input.Text()]++
	}
}
