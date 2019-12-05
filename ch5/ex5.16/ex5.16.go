package main

import (
	"bytes"
	"fmt"
)

//练习5.16：编写多参数版本的strings.Join。

func main() {
	fmt.Println(Join("/", "hi", "there"))
	fmt.Println(Join("/"))
}

func Join(sep string, strs ...string) string {
	if len(strs) == 0 {
		return ""
	}
	b := bytes.Buffer{}
	for _, s := range strs[:len(strs)-1] {
		b.WriteString(s)
		b.WriteString(sep)
	}
	b.WriteString(strs[len(strs)-1])
	return b.String()
}
