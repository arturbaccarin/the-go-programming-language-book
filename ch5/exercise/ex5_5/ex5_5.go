package main

import (
	"fmt"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

// Exercise 5.5: Implement countWordsAndImages.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.TextNode {
		words += len(strings.Fields(n.Data))
	} else if n.Type == html.ElementNode && n.Data == "img" {
		images++
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func main() {
	url := "https://example.com"
	words, images, err := CountWordsAndImages(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Words: %d, Images: %d\n", words, images)
}
