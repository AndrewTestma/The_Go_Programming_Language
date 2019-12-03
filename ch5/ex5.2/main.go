package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"os"
)

//练习 5.2：编写函数，记录在HTML树中出现的同名元素的次数。

func tagFreq(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)
	z := html.NewTokenizer(os.Stdin)
	var err error
	for {
		type_ := z.Next()
		if type_ == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}

func main() {
	freq, err := tagFreq(os.Stdin)
	if err != nil {
		logrus.Fatal(err)
	}
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}
