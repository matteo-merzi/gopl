package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"

	"gopl.io/ch4/github"
)

var instructions = `usage:
search QUERY
[read|edit|close|open] OWNER REPO ISSUE_NUMBER
`

func printUsage() {
	fmt.Fprintf(os.Stderr, instructions)
	os.Exit(1)
}

func main() {
	if len(os.Args) < 2 {
		printUsage()
	}
	cmd := os.Args[1]
	args := os.Args[2:]
	if (cmd == "search" && len(args) < 1) || (cmd != "search" && len(args) != 3) {
		printUsage()
	}
	owner, repo, issueNumber := args[0], args[1], args[2]
	switch cmd {
	case "search":
		search(args)
	case "read":
		read(owner, repo, issueNumber)
	case "edit":
		edit(owner, repo, issueNumber)
	case "close":
		close(owner, repo, issueNumber)
	case "open":
		open(owner, repo, issueNumber)
	}
}

func search(query []string) {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	for _, item := range result.Items {
		format := "#%-5d %9.9s %.55s\n"
		fmt.Printf(format, item.Number, item.User.Login, item.Title)
	}
}

func read(owner, repo, number string) {
	issue, err := getIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}
	body := issue.Body
	if body == "" {
		body = "<empty>\n"
	}
	fmt.Printf("repo: %s/%s\nnumber: %s\nuser: %s\ntitle: %s\n\n%s",
		owner, repo, number, issue.User.Login, issue.Title, body)
}

func edit(owner, repo, number string) {
	editor := os.Getenv("EDITOR")
	if editor == "" {
		editor = "vim"
	}
	editorPath, err := exec.LookPath(editor)
	if err != nil {
		log.Fatal(err)
	}
	tempfile, err := ioutil.TempFile("", "issue_crud")
	if err != nil {
		log.Fatal(err)
	}
	defer tempfile.Close()
	defer os.Remove(tempfile.Name())

	issue, err := getIssue(owner, repo, number)
	if err != nil {
		log.Fatal(err)
	}

	encoder := json.NewEncoder(tempfile)
	err = encoder.Encode(map[string]string{
		"title": issue.Title,
		"state": issue.State,
		"body":  issue.Body,
	})
	if err != nil {
		log.Fatal(err)
	}

	cmd := &exec.Cmd{
		Path:   editorPath,
		Args:   []string{editor, tempfile.Name()},
		Stdin:  os.Stdin,
		Stdout: os.Stdout,
		Stderr: os.Stderr,
	}
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	tempfile.Seek(0, 0)
	fields := make(map[string]string)
	if err = json.NewDecoder(tempfile).Decode(&fields); err != nil {
		log.Fatal(err)
	}

	_, err = editIssue(owner, repo, number, fields)
	if err != nil {
		log.Fatal(err)
	}
}

func close(owner, repo, number string) {
	_, err := editIssue(owner, repo, number, map[string]string{"state": "closed"})
	if err != nil {
		log.Fatal(err)
	}
}

func open(owner, repo, number string) {
	_, err := editIssue(owner, repo, number, map[string]string{"state": "open"})
	if err != nil {
		log.Fatal(err)
	}
}
