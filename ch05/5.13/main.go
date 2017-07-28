package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"

	"gopl.io/ch5/links"
)

func breadthFirst(f func(item string) []string, worklist []string) {
	seen := make(map[string]bool)
	for len(worklist) > 0 {
		items := worklist
		worklist = nil
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				worklist = append(worklist, f(item)...)
			}
		}
	}
}

var originHost string

func save(rawurl string) error {
	url, err := url.Parse(rawurl)
	if err != nil {
		return nil
	}
	if originHost == "" {
		originHost = url.Host
	}
	if originHost != url.Host {
		return nil
	}
	dir := url.Host
	var fileName string
	if filepath.Ext(fileName) == "" {
		dir = filepath.Join(dir, url.Path)
		fileName = filepath.Join(dir, "index.html")
	} else {
		dir = filepath.Join(dir, filepath.Dir(url.Path))
		fileName = url.Path
	}
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		return err
	}
	resp, err := http.Get(rawurl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	// Check for delayed write errors, as mentioned at the end of section 5.8.
	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func crawl(url string) []string {
	fmt.Println(url)
	err := save(url)
	if err != nil {
		log.Printf(`can't cache "%s": %s`, url, err)
	}
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

func main() {
	breadthFirst(crawl, os.Args[1:])
}
