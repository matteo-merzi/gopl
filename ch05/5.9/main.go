// ex5.9 expands shell-style variable references on stdin.
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func expand(s string, f func(string) string) string {
	if strings.HasPrefix(s, "$") {
		return f(s[1:])
	}
	return s
}

func main() {
	subs := make(map[string]string, 0)
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			pair := strings.Split(arg, "=")
			if len(pair) != 2 {
				log.Fatalln("parameters must be in this form KEY=VALE")
			}
			key, value := pair[0], pair[1]
			subs[key] = value
		}
	}

	fmt.Println(subs)

	f := func(s string) string {
		v, ok := subs[s]
		if ok {
			return v
		}
		return s
	}

	file, err := os.Open("./text")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		fmt.Printf("%s ", expand(word, f))
	}
}
