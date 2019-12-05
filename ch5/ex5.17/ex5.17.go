package ex5_17

import "golang.org/x/net/html"

//练习5.17：编写多参数版本的ElementsByTagName，函数接收一个HTML结点树以及任意数量的标签名，返回与这些标签名匹配的所有元素。下面给出了2个例子：

func ElementByTagName(n *html.Node, name ...string) []*html.Node {
	nodes := make([]*html.Node, 0)
	keep := make(map[string]bool, len(name))
	for _, t := range name {
		keep[t] = true
	}
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		_, ok := keep[n.Data]
		if ok {
			nodes = append(nodes, n)
		}
		return true
	}
	forEachElement(n, pre, nil)
	return nodes
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
