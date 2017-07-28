package main

import (
	"fmt"
)

func main() {
	s := []string{"a", "a", "b", "b", "c", "d", "d", "e", "e"}
	s = removeDupluicate(s)
	fmt.Println(s)
}

func removeDupluicate(s []string) []string {
	i := 0
	for j, v := range s {
		if j >= len(s)-1 {
			s[i] = v
			i++
			break
		}
		if v != s[j+1] {
			s[i] = v
			i++
		}
	}

	return s[:i]
}
