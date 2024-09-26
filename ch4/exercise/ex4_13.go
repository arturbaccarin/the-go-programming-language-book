package exercise

/*
Exercise 4.13: The JSON-based  web  service  of the Open Movie Databas e lets you search
https://omdbapi.com/ for a movie by name and downlo ad its poster image. Write a tool
poster that downloads the poster image for the movie named on the command line.
*/

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

const apiKey = "YOUR_API_KEY" // Replace with your OMDB API key

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func fetchMoviePoster(movieName string) (string, error) {
	url := fmt.Sprintf("https://www.omdbapi.com/?apikey=%s&t=%s", apiKey, movieName)
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("failed to fetch movie data")
	}

	var movie Movie
	if err := json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return "", err
	}

	if movie.Poster == "" {
		return "", fmt.Errorf("no poster found for movie: %s", movieName)
	}

	return movie.Poster, nil
}

func downloadPoster(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	return err
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: poster <movie name>")
		return
	}

	movieName := strings.Join(os.Args[1:], " ")
	posterURL, err := fetchMoviePoster(movieName)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	filename := "poster.jpg"
	if err := downloadPoster(posterURL, filename); err != nil {
		fmt.Println("Error downloading poster:", err)
		return
	}

	fmt.Println("Poster downloaded successfully as", filename)
}
