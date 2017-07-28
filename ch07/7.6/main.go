package main

import (
	"flag"
	"fmt"

	"github.com/matteo-merzi/gopl/ch07/7.6/tempflag"
)

var temp = tempflag.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Println(*temp)
}
