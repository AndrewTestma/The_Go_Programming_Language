package ex7_2

import (
	"bytes"
	"testing"
)

func TestCountingWriter(t *testing.T) {
	b := &bytes.Buffer{}
	c, n := CountWriter(b)
	data := []byte("hello world")
	c.Write(data)
	if *n != int64(len(data)) {
		t.Logf("%d != %d", n, len(data))
		t.Fail()
	} else {
		t.Logf("%d == %d", *n, len(data))
	}
}
