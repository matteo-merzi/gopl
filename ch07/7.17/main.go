// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 214.
//!+

// Xmlselect prints the text of selected elements of an XML document.
package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type nodeStack struct {
	nodes []node
}

func (l *nodeStack) Push(n node) {
	l.nodes = append(l.nodes, n)
}

func (l *nodeStack) Pop() {
	l.nodes = l.nodes[:len(l.nodes)-1]
}

func (l *nodeStack) Head() *node {
	return &l.nodes[0]
}

func (l *nodeStack) Tail() *nodeStack {
	return &nodeStack{l.nodes[1:]}
}

func (l *nodeStack) Size() int {
	return len(l.nodes)
}

func (l *nodeStack) String() string {
	var b bytes.Buffer
	for i, n := range l.nodes {
		if i > 0 {
			b.WriteString(" ")
		}
		b.WriteString(n.name)
		for name, value := range n.attributes {
			b.WriteString(fmt.Sprintf(`|%s="%s"`, name, value))
		}
	}
	return b.String()
}

type node struct {
	name       string
	attributes map[string]string
}

func (n *node) IsSubset(x *node) bool {
	if n.name != x.name {
		return false
	}

	if len(n.attributes) > len(x.attributes) {
		return false
	}

	for key, value := range n.attributes {
		v, ok := x.attributes[key]
		if !ok || value != v {
			return false
		}
	}

	return true
}

type attribute struct {
	name  string
	value string
}

func parseInput(args []string) (*nodeStack, error) {
	nodes := make([]node, 0)
	for _, arg := range args {
		attributes := make(map[string]string)
		fields := strings.Split(arg, "|")
		node := node{}
		for i, f := range fields {
			if i == 0 {
				node.name = f
				continue
			}
			fx := strings.Split(f, "=")
			if len(fx) != 2 {
				return nil, fmt.Errorf("wrong input: %s", f)
			}
			name := fx[0]
			val := fx[1]
			val = val[1 : len(val)-1]
			attributes[name] = val
		}
		node.attributes = attributes
		nodes = append(nodes, node)
	}
	return &nodeStack{nodes}, nil
}

// containsAll reports whether x contains the elements of y, in order.
func containsAll(x, y *nodeStack) bool {
	for y.Size() <= x.Size() {
		if y.Size() == 0 {
			return true
		}

		yH := y.Head()
		xH := x.Head()
		if yH.IsSubset(xH) {
			y = y.Tail()
		}
		x = x.Tail()
	}
	return false
}

func xmlSelect(w io.Writer, r io.Reader, selector *nodeStack) error {
	dec := xml.NewDecoder(r)
	stack := &nodeStack{} // stack of node element
	for {
		tok, err := dec.Token()
		if err == io.EOF {
			break
		} else if err != nil {
			return fmt.Errorf("xmlselect: %v", err)
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			attributes := make(map[string]string)
			for _, attr := range tok.Attr {
				attributes[attr.Name.Local] = attr.Value
			}
			stack.Push(node{name: tok.Name.Local, attributes: attributes}) // push
		case xml.EndElement:
			stack.Pop() // pop
		case xml.CharData:
			if containsAll(stack, selector) {
				fmt.Fprintf(w, "%s: %s\n", stack.String(), tok)
			}
		}
	}
	return nil
}

func main() {
	selector, err := parseInput(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	err = xmlSelect(os.Stdout, os.Stdin, selector)
	if err != nil {
		log.Fatal(err)
	}
}

//!-
