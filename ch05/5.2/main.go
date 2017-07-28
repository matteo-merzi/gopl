package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	m := make(map[string]int)
	mapTags(m, doc)
	for k, v := range m {
		fmt.Printf("Tag: %s\tCount: %d\n", k, v)
	}
}

func mapTags(m map[string]int, n *html.Node) {
	if n.Type == html.ElementNode {
		switch n.Data {
		case "a":
			m["a"]++
		case "p":
			m["p"]++
		case "div":
			m["div"]++
		case "br":
			m["br"]++
		case "h1":
			m["h1"]++
		case "h2":
			m["h2"]++
		case "h3":
			m["h3"]++
		default:
			m["other"]++
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		mapTags(m, c)
	}
}
