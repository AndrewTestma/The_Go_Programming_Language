package main

// 练习 7.9：使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。

import (
	ex7_8 "The_Go_Programming_Language/ch7/ex7.8"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"sort"
)

var people = []ex7_8.Person{
	{"Alice", 20},
	{"Bob", 12},
	{"Bob", 20},
	{"Alice", 12},
}

var html = template.Must(template.New("people").Parse(`
<html>
<body>
<table>
	<tr>
		<th><a href="?sort=name">name</a></th>
		<th><a href="?sort=age">age</a></th>
	</tr>
{{range .}}
	<tr>
		<td>{{.Name}}</td>
		<td>{{.Age}}</td>
	</td>
{{end}}
</body>
</html>
`))

func main() {
	c := ex7_8.NewByColumns(people, 2)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.FormValue("sort") {
		case "age":
			c.Select(c.LessAge)
		case "name":
			c.Select(c.LessName)
		}
		sort.Sort(c)
		fmt.Println(people)
		err := html.Execute(w, people)
		if err != nil {
			log.Printf("template error: %s", err)
		}
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
}
