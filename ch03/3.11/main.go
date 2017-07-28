package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	fmt.Printf("Formatted with comma: %s\n", comma("+1234567890.1234"))
}

func comma(s string) string {
	buf := &bytes.Buffer{}
	mantissaStart := 0

	if s[0] == '+' || s[0] == '-' {
		buf.WriteByte(s[0])
		mantissaStart = 1
	}
	mantissaEnd := strings.Index(s, ".")
	if mantissaEnd == -1 {
		mantissaEnd = len(s)
	}
	mantissa := s[mantissaStart:mantissaEnd]

	start := len(mantissa) % 3
	if start == 0 {
		start = 3
	}
	buf.WriteString(mantissa[:start])

	for i := start; i < len(mantissa); i += 3 {
		buf.WriteByte(',')
		buf.WriteString(mantissa[i : i+3])
	}

	buf.WriteString(s[mantissaEnd:])

	return buf.String()
}
