package main

import (
	"fmt"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Fprintln(os.Stderr, "usage: main HTML_FILE ID")
	}
	filename := os.Args[1]
	id := os.Args[2]
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	doc, err := html.Parse(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v\n", ElementByID(doc, id))
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}
		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}
		return true
	}

	return forEachNode(doc, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		stop := !pre(n)
		if stop {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		stop := forEachNode(c, pre, post)
		if stop != nil {
			return stop
		}
	}

	if post != nil {
		stop := !post(n)
		if stop {
			return n
		}
	}

	return nil
}
