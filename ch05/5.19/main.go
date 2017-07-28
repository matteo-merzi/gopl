package main

import "fmt"

func falsePanic() (s string) {
	type falsePanic struct{}

	defer func() {
		switch p := recover(); p {
		case falsePanic{}:
			s = "false panic"
		default:
			panic(p)
		}
	}()
	panic(falsePanic{})
}

func main() {
	fmt.Printf("result: %s\n", falsePanic())
}
