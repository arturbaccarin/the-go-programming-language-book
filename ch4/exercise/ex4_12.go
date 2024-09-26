package exercise

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

/*
Exercise 4.12: The popular web comic xkcd has a JSON interface. For example, a request to
https://xkcd.com/571/info.0.json produces a detailed description of comic 571, one of
many favorites.  Download each URL (once!) and build an ofﬂine index.  Write a tool xkcd
that, using this index, prints the URL and transcript of each comic that matches a search term
provided on the command line.
*/

type Comic struct {
	Month      string `json:"month"`
	Num        int    `json:"num"`
	Link       string `json:"link"`
	Year       string `json:"year"`
	News       string `json:"news"`
	SafeTitle  string `json:"safe_title"`
	Transcript string `json:"transcript"`
	Alt        string `json:"alt"`
	Img        string `json:"img"`
	Title      string `json:"title"`
	Day        string `json:"day"`
}

var index = map[string]*Comic{}

func Find(number string) *Comic {
	comic, ok := index[number]
	if ok {
		return comic
	}

	comic, err := requestComic(number)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	index[number] = comic

	return comic
}

func requestComic(number string) (*Comic, error) {
	url := fmt.Sprintf("https://xkcd.com/%s/info.0.json", number)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Unmarshal JSON data into the Comic struct
	var comic Comic
	if err := json.Unmarshal(body, &comic); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return &comic, nil
}
