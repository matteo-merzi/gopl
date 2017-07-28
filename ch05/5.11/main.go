package main

import (
	"fmt"
	"log"
	"strings"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	order, err := topoSort(prereqs)
	if err != nil {
		log.Fatalln(err)
	}
	for i, course := range order {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func indexOf(s string, sx []string) (int, error) {
	for i, v := range sx {
		if v == s {
			return i, nil
		}
	}
	return -1, fmt.Errorf("Element not found")
}

func topoSort(m map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)
	var visitAll func(items []string, parents []string)

	visitAll = func(items []string, parents []string) {
		for _, v := range items {
			isResolved, seen := resolved[v]
			if !isResolved && seen {
				start, _ := indexOf(v, parents)
				err = fmt.Errorf("cycle: %s", strings.Join(append(parents[start:], v), " -> "))
			}
			if !seen {
				resolved[v] = false
				visitAll(m[v], append(parents, v))
				resolved[v] = true
				order = append(order, v)
			}
		}
	}

	for k := range m {
		if err != nil {
			return nil, err
		}
		visitAll([]string{k}, nil)
	}
	return order, nil
}
