package ex7_5

import (
	"bytes"
	"strings"
	"testing"
)

func TestLimitReader(t *testing.T) {
	s := "hello andrew"
	b := &bytes.Buffer{}
	r := LimitReader(strings.NewReader(s), 5)
	n, _ := b.ReadFrom(r)
	if n != 5 {
		t.Logf("n=%d", n)
		t.Fail()
	}
	if b.String() == "hello" {
		t.Logf(`"%s" == "%s"`, b.String(), s[:5])
	}
}
