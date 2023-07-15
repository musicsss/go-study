package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

const OMDBAPIURL = "https://omdbapi.com/"

type Movie struct {
	Title  string `json:"Title"`
	Poster string `json:"Poster"`
}

func GetMoviePoster(movieName string) (string, error) {
	query := url.QueryEscape(movieName)
	resp, err := http.Get(OMDBAPIURL + "?t=" + query)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	var movie Movie

	err = json.NewDecoder(resp.Body).Decode(&movie)
	if err != nil {
		return "", err
	}
	log.Println(movie)
	return movie.Poster, nil
}
func DownloadPoster(posterURL string, filename string) error {
	log.Println(posterURL)
	resp, err := http.Get(posterURL)
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
	if err != nil {
		return err
	}
	return nil
}
func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide a movie name.")
		return
	}
	movieName := strings.Join(os.Args[1:], " ")
	posterURL, err := GetMoviePoster(movieName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("1111111")
	err = DownloadPoster(posterURL, "poster.jpg")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Poster downloaded successfully.")
}
