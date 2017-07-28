package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}
	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.TextNode:
			words += wordsCount(c.Data)
		case html.ElementNode:
			if c.Data == "img" {
				images++
			}
		}
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func wordsCount(s string) int {
	scan := bufio.NewScanner(strings.NewReader(s))
	scan.Split(bufio.ScanWords)
	words := 0
	for scan.Scan() {
		words++
	}
	return words
}

func main() {
	if len(os.Args) != 2 {
		log.Fatalf("Url has to been pass as a parameter\n")
	}
	words, images, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("Words count: %d\n", words)
	fmt.Printf("Images count: %d\n", images)
}
