package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
	"os"
	"strings"
)

//练习 5.7：完善startElement和endElement函数，使其成为通用的HTML输出器。要求：输出注释结点，文本结点以及每个元素的属性（< a href='...'>）。使用简略格式输出没有孩子结点的元素（即用<img/>代替<img></img>）。编写测试，验证程序输出的格式正确。（详见11章）
var depth int

type PrettyPrinter struct {
	w   io.Writer
	err error
}

func NewPrettyPrinter() PrettyPrinter {
	return PrettyPrinter{}
}

func (pp PrettyPrinter) Pretty(w io.Writer, n *html.Node) error {
	pp.w = w
	pp.err = nil
	pp.forEachNode(n, pp.start, pp.end)
	return pp.Err()
}

func (pp PrettyPrinter) Err() error {
	return pp.err
}

func (pp PrettyPrinter) forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	if pp.Err() != nil {
		return
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		pp.forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
	if pp.Err() != nil {
		return
	}
}

func (pp PrettyPrinter) printf(format string, args ...interface{}) {
	_, err := fmt.Fprintf(pp.w, format, args...)
	pp.err = err
}

func (pp PrettyPrinter) startElement(n *html.Node) {
	end := ">"
	if n.FirstChild == nil {
		end = "/>"
	}
	attrs := make([]string, 0, len(n.Attr))
	for _, a := range n.Attr {
		attrs = append(attrs, fmt.Sprintf(`%s="%s"`, a.Key, a.Val))
	}
	attrStr := ""
	if len(n.Attr) > 0 {
		attrStr = " " + strings.Join(attrs, " ")
	}
	name := n.Data
	pp.printf("%*s<%s%s%s\n", depth*2, "", name, attrStr, end)
	depth++
}
func (pp PrettyPrinter) endElement(n *html.Node) {
	depth--
	if n.FirstChild == nil {
		return
	}
	pp.printf("%*s</%s>\n", depth*2, "", n.Data)
}

func (pp PrettyPrinter) startText(n *html.Node) {
	text := strings.TrimSpace(n.Data)
	if len(text) == 0 {
		return
	}
	pp.printf("%*s%s\n", depth*2, "", n.Data)
}

func (pp PrettyPrinter) startComment(n *html.Node) {
	pp.printf("<!--%s-->\n", n.Data)
}

func (pp PrettyPrinter) start(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.startElement(n)
	case html.TextNode:
		pp.startText(n)
	case html.CommentNode:
		pp.startComment(n)
	}
}

func (pp PrettyPrinter) end(n *html.Node) {
	switch n.Type {
	case html.ElementNode:
		pp.endElement(n)
	}
}
func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}
	pp := NewPrettyPrinter()
	pp.Pretty(os.Stdout, doc)
	return nil
}
