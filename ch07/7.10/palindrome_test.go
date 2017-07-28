package palindrome

import (
	"sort"
	"testing"
)

func TestIsPalidrome(t *testing.T) {
	s := []string{"*", "**", "***", "**", "*"}
	if !IsPalindrome(sort.StringSlice(s)) {
		t.Fail()
	}
}

func TestIsNotPalidrome(t *testing.T) {
	s := []string{"*", "**", "***", "**", "**", "*"}
	if IsPalindrome(sort.StringSlice(s)) {
		t.Fail()
	}
}
