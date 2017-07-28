package main

import (
	"bytes"
	"fmt"
)

func main() {
	fmt.Printf("Formatted with comma: %s\n", comma("123456789"))
}

func comma(s string) string {
	start := len(s) % 3
	buf := &bytes.Buffer{}

	if start == 0 {
		start = 3
	}
	buf.WriteString(s[:start])

	for i := start; i < len(s); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(s[i : i+3])
	}

	return buf.String()
}
