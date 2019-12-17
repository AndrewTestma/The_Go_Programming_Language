package ex7_10

import (
	"sort"
	"testing"
)

func TestIsPalindrome_oddNumElements(t *testing.T) {
	ints := []int{1, 2, 3, 4, 3, 2, 1}
	if !IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}
func TestPalindrome_evenNumElements(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 1}
	if !IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}

func TestPalindrome_negative(t *testing.T) {
	ints := []int{1, 2, 3, 3, 2, 2}
	if IsPalindrome(sort.IntSlice(ints)) {
		t.Fail()
	}
}
