package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func myJoin(sep string, xs ...string) string {
	return strings.Join(xs, sep)
}

func main() {
	if len(os.Args) == 0 {
		log.Fatalln("USAGE: main SEP ...ARGS")
	}

	sep := os.Args[1]
	xs := make([]string, 0)
	for _, v := range os.Args[2:] {
		xs = append(xs, v)
	}

	fmt.Printf("Jointed string: %s\n", myJoin(sep, xs...))
}
