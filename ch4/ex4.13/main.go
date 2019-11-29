package main

import (
	"encoding/json"
	"net/http"
)

const APIURL = "http://omdbapi.com/"

type Movie struct {
	Titie  string
	Year   string
	Poster string
}

func (m Movie) posterFileName(title string) {

}

func getMovie(title string) (*Movie, error) {
	resp, err := http.Get(APIURL + "?t=" + title)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		_ = resp.Body.Close()
		return nil, err
	}
	var movie Movie
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		_ = resp.Body.Close()
		return nil, err
	}
	return &movie, nil
}

func main() {

}
