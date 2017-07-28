package reverse

import (
	"fmt"
	"unicode/utf8"
)

func rev(b []byte) {
	size := len(b)
	for i := 0; i < len(b)/2; i++ {
		b[i], b[size-1-i] = b[size-1-i], b[i]
	}
}

// Reverse all the runes, and then the entire slice. The runes' bytes end up in
// the right order.
func revUTF8(b []byte) []byte {
	fmt.Printf("%q\n", string(b))
	for i := 0; i < len(b); {
		r, size := utf8.DecodeRune(b[i:])
		fmt.Printf("Actual rune: %q\n", string(r))
		rev(b[i : i+size])
		i += size
	}
	rev(b)
	return b
}
