package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Printf("%s\n", s)
	elapsed := time.Since(start)
	fmt.Printf("for cycle variant took %s\n", elapsed)

	start = time.Now()
	fmt.Printf("%s\n", strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start)
	fmt.Printf("string.Join variant took %s\n", elapsed)

}
