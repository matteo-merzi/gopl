package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"

	eval "github.com/matteo-merzi/gopl/ch07/7.14"
)

func main() {
	expStr := flag.String("expression", "", "the expression to evaluate")
	envStr := flag.String("variables", "", "the variables values. ex. x=1 y=2 z=...")
	flag.Parse()

	env := eval.Env{}
	assignments := strings.Split(*envStr, " ")
	for _, a := range assignments {
		fields := strings.Split(a, "=")
		if len(fields) != 2 {
			log.Fatalf("incorrect variable assignment: %v", fields)
		}
		name, valueStr := fields[0], fields[1]
		fmt.Printf("name: %s, value: %s\n", name, valueStr)
		val, err := strconv.ParseFloat(valueStr, 64)
		if err != nil {
			log.Fatalf("unparsable value for %s", name)
		}
		env[eval.Var(name)] = val
	}

	exp, err := eval.Parse(*expStr)
	if err != nil {
		log.Fatalf("unparsable expression: %s", *expStr)
	}
	fmt.Printf("expression: %s\n", exp.String())

	vars := make(map[eval.Var]bool)
	err = exp.Check(vars)
	if err != nil {
		for _, v := range vars {
			fmt.Println(v)
		}
		log.Fatalf("error checking expression: %s", *expStr)
	}
	fmt.Printf("Result: %f\n", exp.Eval(env))
}
