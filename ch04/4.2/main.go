package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

func main() {
	var methodStr string
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&methodStr, "m", "256", "sha method")
	mySet.Parse(os.Args[1:])

	var input string
	fmt.Println("Enter text:")
	for {
		fmt.Scan(&input)
		if input == "" {
			continue
		}
		value := []byte(input)
		switch methodStr {
		case "384":
			fmt.Printf("SHA384 of %s: %x\n", input, sha512.Sum384(value))
		case "512":
			fmt.Printf("SHA512 of %s: %x\n", input, sha512.Sum512(value))
		default:
			fmt.Printf("SHA256 of %s: %x\n", input, sha256.Sum256(value))
		}
	}
}
