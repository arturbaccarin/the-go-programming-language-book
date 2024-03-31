package exercise

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

/* test
Exercise 1.7 - The function call io.Copy(dst, src) reads from src and writes to dst. Use it
instead of ioutil.ReadAll to  copy the response body to os.Stdout without requiring a
buffer large enough to hold the entire stream. Be sure to check the error result of io.Copy.
*/

/*
Exercise 1.8 - Modify fetch to add the preﬁx http:// to each argument URL if it is missing.
You might want to use strings.HasPrefix.
*/

/*
Exercise 1.9 - Modify fetch to also print the HTTP status code, found in resp.Status.
*/

const httpPrefix = "http://"

func Exercise1_7_8_9() {
	for _, url := range os.Args[1:] {
		if !strings.HasPrefix(url, httpPrefix) {
			url = httpPrefix + url
		}

		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Printf("Status code: %v", resp.StatusCode)

		_, err = io.Copy(os.Stdout, resp.Body)
		resp.Body.Close()
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
