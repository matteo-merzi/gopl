package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()

	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("Less then a month old: ")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 < 30 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("Less then a year old: ")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 >= 30 && now.Sub(item.CreatedAt).Hours()/24 < 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
	fmt.Println("More then a year old: ")
	for _, item := range result.Items {
		if now.Sub(item.CreatedAt).Hours()/24 >= 365 {
			fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
		}
	}
}
