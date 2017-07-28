package main

import (
	"fmt"
)

func main() {
	s := []int{0, 1, 2, 3, 4, 5}
	s = rotate(s, 3)
	fmt.Println(s)
}

// 0 1 2 3 4 5 START
// 1 2 3 4 5 0 rotate 1
// 2 3 4 5 0 1 rotate 2

func rotate(s []int, n int) []int {
	return append(s[n:len(s)], s[:n]...)
}
