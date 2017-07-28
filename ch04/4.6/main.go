package main

import (
	"fmt"
	"unicode"
)

func main() {
	b := []byte("hello\r  \tworld")
	fmt.Printf("Before: %q\n", string(b))
	fmt.Printf("After: %q\n", string(squashUnicodeSpace(b)))
}

func squashUnicodeSpace(b []byte) []byte {
	out := b[:0]
	for i, c := range b {
		if unicode.IsSpace(rune(c)) {
			if i > 0 && unicode.IsSpace(rune(b[i-1])) {
				continue
			} else {
				out = append(out, ' ')
			}
		} else {
			out = append(out, c)
		}
	}
	return out
}
