package main

import (
	"The_Go_Programming_Language/ch4/github"
	"fmt"
	"github.com/Sirupsen/logrus"
	"html/template"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var issueListTemplate = template.Must(template.New("issueList").Parse(`
<h1>{{.Issues | len}} issues</h1>
<table>
<tr style='text-align: left'>
<th>#</th>
<th>State</th>
<th>User</th>
<th>Title</th>
</tr>
{{range .Issues}}
<tr>
	<td><a href='{{.CacheURL}}'>{{.Number}}</td>
	<td>{{.State}}</td>
	<td><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></td>
	<td><a href='{{.CacheURL}}'>{{.Title}}</a></td>
</tr>
{{end}}
</table>
`))

var issueTemplate = template.Must(template.New("issue").Parse(`
<h1>{{.Title}}</h1>
<dl>
	<dt>user</dt>
	<dd><a href='{{.User.HTMLURL}}'>{{.User.Login}}</a></dd>
	<dt>state</dt>
	<dd>{{.State}}</dd>
</dl>
<p>{{.Body}}</p>
`))

type IssueCache struct {
	Issues         []github.Issue
	IssuesByNumber map[int]github.Issue
}

func NewIssueCache(owner, repo string) (ic IssueCache, err error) {
	issues, err := github.GetIssues(owner, repo)
	if err != nil {
		return
	}
	ic.Issues = issues
	ic.IssuesByNumber = make(map[int]github.Issue, len(issues))
	for _, issue := range issues {
		ic.IssuesByNumber[issue.Number] = issue
	}
	return
}
func logNonNil(v interface{}) {
	if v != nil {
		logrus.Print(v)
	}
}

func (ic IssueCache) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.SplitN(r.URL.Path, "/", -1)
	if len(pathParts) < 3 || pathParts[2] == "" {
		logNonNil(issueListTemplate.Execute(w, ic))
		return
	}
	numstr := pathParts[2]
	num, err := strconv.Atoi(numstr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err := w.Write([]byte(fmt.Sprintf("Issue number isn't number: '%s'", numstr)))
		if err != nil {
			logrus.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	issue, ok := ic.IssuesByNumber[num]
	if !ok {
		w.WriteHeader(http.StatusNotFound)
		_, err := w.Write([]byte(fmt.Sprintf("No issue '%d'", num)))
		if err != nil {
			logrus.Printf("Error writing response for %s: %s", r, err)
		}
		return
	}
	logNonNil(issueTemplate.Execute(w, issue))
}
func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: githubserver OWNER REPO")
		os.Exit(1)
	}
	owner := os.Args[1]
	repo := os.Args[2]

	issueCache, err := NewIssueCache(owner, repo)
	if err != nil {
		logrus.Fatal(err)
	}
	http.Handle("/", issueCache)
	logrus.Fatal(http.ListenAndServe(":8080", nil))
}
