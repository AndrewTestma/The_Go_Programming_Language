package main

import (
	"fmt"
	"github.com/Sirupsen/logrus"
	"golang.org/x/net/html"
	"net/http"
	"os"
)

//练习 5.8：修改pre和post函数，使其返回布尔类型的返回值。返回false时，中止forEachNoded的遍历。使用修改后的代码编写ElementByID函数，根据用户输入的id查找第一个拥有该id元素的HTML元素，查找成功后，停止遍历。

func ElementById(n *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
		return true
	}
	return forEachElement(n, pre, nil)
}

func forEachElement(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	u := make([]*html.Node, 0)
	u = append(u, n)
	for len(u) > 0 {
		n = u[0]
		u = u[1:]
		if pre != nil {
			if !pre(n) {
				return n
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			u = append(u, c)
		}
		if post != nil {
			if !post(n) {
				return n
			}
		}
	}
	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintf(os.Stderr, "usage: ex5.8 HTML_URL ID")
	}
	resp, err := http.Get(os.Args[1])
	if err != nil {
		logrus.Fatal(err)
	}
	doc, err := html.Parse(resp.Body)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("%+v\n", ElementById(doc, os.Args[2]))
}
