package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	visit(doc)
}

func visit(n *html.Node) {
	if n.Data == "script" || n.Data == "style" {
		return
	}

	if n.Type == html.TextNode {
		s := strings.TrimSpace(n.Data)
		if len(s) > 0 {
			fmt.Printf("%s\n", s)
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		visit(c)
	}
}
