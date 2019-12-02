package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/Sirupsen/logrus"
	"github.com/metal3d/go-slugify"
	"net/http"
	"os"
	"path/filepath"
)

const APIURL = "http://omdbapi.com/"

type Movie struct {
	Titie  string
	Year   string
	Poster string
}

func (m Movie) posterFileName() string {
	ext := filepath.Ext(m.Titie)
	title := slugify.Marshal(m.Titie)
	return fmt.Sprintf("%s_(%s)%s", title, m.Year, ext)
}

func getMovie(title string) (movie Movie, err error) {
	//url_ := fmt.Sprintf("%st=%s",APIURL,url.QueryEscape(title))
	resp, err := http.Get(APIURL + "?t=" + title)
	if err != nil {
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("%d response from ", resp.StatusCode)
		return
	}
	if err = json.NewDecoder(resp.Body).Decode(&movie); err != nil {
		return
	}
	return
}
func (m Movie) writePoster() error {
	url_ := m.Poster
	resp, err := http.Get(url_)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("%d response from %s", resp.StatusCode, url_)
	}
	file, err := os.Create(m.posterFileName())
	if err != nil {
		return err
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	_, err = writer.ReadFrom(resp.Body)
	if err != nil {
		return err
	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}

func main() {
	if len(os.Args) != 2 {
		fmt.Fprintln(os.Stderr, "usage: poster MoVIE_TITLE")
		os.Exit(1)
	}
	title := os.Args[1]
	movie, err := getMovie(title)
	if err != nil {
		logrus.Fatal(err)
	}
	if zero := new(Movie); movie == *zero {
		fmt.Fprintf(os.Stderr, "No results for '%s'\n", title)
		os.Exit(2)
	}
	err = movie.writePoster()
	if err != nil {
		logrus.Fatal(err)
	}
}
