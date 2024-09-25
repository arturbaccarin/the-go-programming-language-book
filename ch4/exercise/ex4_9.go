package exercise

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
Exercise 4.9: Write a program wordfreq to report the frequency of each word in an input text
ﬁle. Call input.Split(bufio.ScanWords) before the ﬁrst call to Scan to break the input into
words instead of lines.
*/
func main() {
	// Open the input file
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a scanner to read the file
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	// Create a map to store word frequencies
	wordFreq := make(map[string]int)

	// Scan the file and count word frequencies
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		wordFreq[word]++
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	}

	// Print the word frequencies
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}
