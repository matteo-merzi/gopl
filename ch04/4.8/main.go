package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
	"unicode/utf8"
)

const (
	CHAR_IS_SPACE = iota
	CHAR_IS_SYMBOL
	CHAR_IS_MARK
	CHAR_IS_DIGIT
	CHAR_IS_PUNCT
	CHAR_IS_LETTER
	CHAR_IS_CONTROL
	CHAR_IS_GRAPHIC
)

func main() {
	counts := make(map[rune]int)    // counts of Unicode characters
	kinds := make(map[int]int)      // counts the kinds of characters used
	var utflen [utf8.UTFMax + 1]int // count of lengths of UTF-8 encodings
	invalid := 0                    // count of invalid UTF-8 characters

	in := bufio.NewReader(os.Stdin)

	// for test purpouse the program will only process the first 50 characters
	limit := 50

	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		// r could be letter, digit, and so on (symbol...)
		switch {
		case unicode.IsSpace(r):
			kinds[CHAR_IS_SPACE]++
		case unicode.IsSymbol(r):
			kinds[CHAR_IS_SYMBOL]++
		case unicode.IsMark(r):
			kinds[CHAR_IS_MARK]++
		case unicode.IsDigit(r):
			kinds[CHAR_IS_DIGIT]++
		case unicode.IsPunct(r):
			kinds[CHAR_IS_PUNCT]++
		case unicode.IsLetter(r):
			kinds[CHAR_IS_LETTER]++
		case unicode.IsControl(r):
			kinds[CHAR_IS_CONTROL]++
		case unicode.IsGraphic(r):
			kinds[CHAR_IS_GRAPHIC]++
		}

		counts[r]++
		utflen[n]++

		limit--
		if limit == 0 {
			break
		}
	}
	fmt.Printf("rune\tcount\n")
	for c, n := range counts {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range utflen {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Print("\nkinds\tcount\n")
	var tname string
	for k, v := range kinds {
		switch k {
		case CHAR_IS_SPACE:
			tname = "space"
		case CHAR_IS_SYMBOL:
			tname = "symbol"
		case CHAR_IS_MARK:
			tname = "mark"
		case CHAR_IS_DIGIT:
			tname = "digit"
		case CHAR_IS_PUNCT:
			tname = "punct"
		case CHAR_IS_LETTER:
			tname = "letter"
		case CHAR_IS_CONTROL:
			tname = "control"
		case CHAR_IS_GRAPHIC:
			tname = "graphic"
		}
		fmt.Printf("%s\t%d\n", tname, v)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}
