package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [8]byte

/*
00000000
00000001
00000010
00000100
00001000
00010000
00100000
01000000
10000000
*/
func init() {
	for i := uint(0); i < 8; i++ {
		pc[i] = byte(1 << i)
		fmt.Printf("%b", pc[i])
	}
}

func main() {
	a := sha256.Sum256([]byte("hello world"))
	b := sha256.Sum256([]byte("hello worlD"))

	printHash(a)
	printHash(b)
	fmt.Println(bitDiff(a, b))
}

func printHash(h [32]byte) {
	for _, v := range h {
		fmt.Printf("%X", v)
	}
	fmt.Println()
}

func bitDiff(a, b [32]byte) int {
	count := 0
	for i := 0; i < 32; i++ {

		// Version A
		// for j := 0; j < 8; j++ {
		// 	if a[i]&pc[j] != b[i]&pc[j] {
		// 		count++
		// 	}
		// }
		// Versions B
		count += popCount(a[i] ^ b[i])
	}
	return count
}

func popCount(b byte) int {
	count := 0
	for ; b != 0; count++ {
		b &= b - 1
	}
	return count
}
