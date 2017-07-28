package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Open("./test.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	words := make(map[string]int)
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		words[word]++
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("wordfreq: %v\n", err)
	}

	for k, v := range words {
		fmt.Printf("Word: %s\tCount: %d\n", k, v)
	}
}
