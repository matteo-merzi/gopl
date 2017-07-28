package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func max(vals ...int) (result int, ok bool) {
	result = math.MinInt64
	if len(vals) == 0 {
		return
	}
	ok = true
	for _, v := range vals {
		if result < v {
			result = v
		}
	}
	return
}

func min(vals ...int) (result int, ok bool) {
	result = math.MaxInt64
	if len(vals) == 0 {
		return
	}
	ok = true
	for _, v := range vals {
		if result > v {
			result = v
		}
	}
	return
}

func main() {
	vals := make([]int, 0)
	for _, v := range os.Args[1:] {
		intV, err := strconv.Atoi(v)
		if err != nil {
			log.Fatalln(err)
		}
		vals = append(vals, intV)
	}

	if result, ok := max(vals...); ok {
		fmt.Printf("Max is: %d\n", result)
	} else {
		fmt.Println("max func received no params")
	}
	if result, ok := min(vals...); ok {
		fmt.Printf("Min is: %d\n", result)
	} else {
		fmt.Println("min func received no params")
	}
}
