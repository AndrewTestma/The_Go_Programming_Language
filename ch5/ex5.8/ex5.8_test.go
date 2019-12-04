package main

import (
	"fmt"
	"golang.org/x/net/html"
	"strings"
	"testing"
)

func TestElementById(t *testing.T) {
	input := `
<html>
<body>
	<p class="something" id = "short">
		<span class="special">hi</span>
	</p>
	<br/>
<body>
</html>
`
	doc, err := html.Parse(strings.NewReader(input))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", ElementById(doc, "short"))
}
