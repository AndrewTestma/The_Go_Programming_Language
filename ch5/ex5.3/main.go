package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
	"io"
	"os"
	"strings"
)

//练习 5.3：编写函数输出所有text结点的内容。注意不要访问<script>和<style>元素,因为这些元素对浏览者是不可见的。

func printTagText(r io.Reader, w io.Writer) error {
	z := html.NewTokenizer(os.Stdin)
	var err error
	stack := make([]string, 20)
Tokenize:
	for {
		switch z.Next() {
		case html.ErrorToken:
			break Tokenize
		case html.StartTagToken:
			b, _ := z.TagName()
			stack = append(stack, string(b))
		case html.TextToken:
			cur := stack[len(stack)-1]
			if cur == "script" || cur == "style" {
				continue
			}
			text := z.Text()
			if len(strings.TrimSpace(string(text))) == 0 {
				continue
			}
			w.Write([]byte(fmt.Sprintf("<%s>", cur)))
			w.Write(text)
			if text[len(text)-1] != '\n' {
				io.WriteString(w, "\n")
			}
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func main() {
	err := printTagText(os.Stdin, os.Stdout)
	if err != nil {
		logrus.Fatal(err)
	}
}
