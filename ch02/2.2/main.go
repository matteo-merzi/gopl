package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/matteo-merzi/gopl/2.2/conv"
)

func convertMeasurement(s string) {
	x, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	m := conv.Meter(x)
	c := conv.Celsius(x)
	k := conv.Kilogram(x)
	fmt.Printf("%s = %s\n", m, conv.MToF(m))
	fmt.Printf("%s = %s\n", c, conv.CToF(c))
	fmt.Printf("%s = %s\n", k, conv.KToP(k))
	fmt.Println("******************")
}

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			convertMeasurement(arg)
		}
	} else {
		scan := bufio.NewScanner(os.Stdin)
		for scan.Scan() {
			convertMeasurement(scan.Text())
		}
	}
}
