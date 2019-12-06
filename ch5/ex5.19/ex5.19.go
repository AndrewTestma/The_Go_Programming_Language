package main

import "fmt"

func main() {
	fmt.Print(weird())
}

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hello"
	}()
	panic("hi")
}
