package ex7_4

import (
	"bytes"
	"golang.org/x/net/html"
	"testing"
)

func TestNewReader(t *testing.T) {
	s := "hi there"
	b := &bytes.Buffer{}
	n, err := b.ReadFrom(NewReader(s))
	if n != int64(len(s)) || err != nil {
		t.Logf("n=%d err=%s", n, err)
		t.Fail()
	} else {
		t.Logf("n=%d", n)
	}
	if b.String() != s {
		t.Logf(`"%s" != "%s"`, b.String(), s)
	}
}
func TestNewReaderWrithHTML(t *testing.T) {
	s := "<html><body><p>hi</p></body></html>"
	_, err := html.Parse(NewReader(s))
	if err != nil {
		t.Log(err)
		t.Fail()
	}
}