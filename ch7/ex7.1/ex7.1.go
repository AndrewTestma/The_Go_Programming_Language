package main

import (
	"bufio"
	"fmt"
	"strings"
)

//练习 7.1：使用来自ByteCounter的思路，实现一个针对对单词和行数的计数器。你会发现bufio.ScanWords非常的有用。

type LineCounter int

func (w *LineCounter) Write(p []byte) (n int, err error) {
	in := bufio.NewScanner(strings.NewReader(string(p)))
	in.Split(bufio.ScanLines)
	for in.Scan() {
		*w++
	}
	return len(p), nil
}

type WordCounter int

func (w *WordCounter) Write(p []byte) (n int, err error) {
	in := bufio.NewScanner(strings.NewReader(string(p)))
	in.Split(bufio.ScanWords)
	for in.Scan() {
		*w++
	}
	return len(p), nil
}
func main() {
	var w WordCounter
	w.Write([]byte("hello world"))
	fmt.Println(w)
	var l LineCounter
	l.Write([]byte("hello\nworld\n\n"))
	fmt.Println(l)
}
