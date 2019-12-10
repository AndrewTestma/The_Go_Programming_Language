package ex7_3

import "testing"

func TestString(t *testing.T) {
	root := &tree{value: 36}
	root = add(root, 21)
	root = add(root, 44)
	if root.String() != "[2 3 4]" {
		t.Log(root)
		t.Fail()
	} else {
		t.Log(root.String())
	}
}
