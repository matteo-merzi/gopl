package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"gopl.io/ch4/github"
)

const IssuesURL = "https://api.github.com/search/issues"
const APIURL = "https://api.github.com"

func get(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("can't get %s: %s", url, resp.Status)
	}
	return resp, nil
}

func getIssue(owner string, repo string, number string) (*github.Issue, error) {
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	resp, err := get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var issue github.Issue
	if err := json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}

func editIssue(owner, repo, number string, fields map[string]string) (*github.Issue, error) {
	buf := &bytes.Buffer{}
	encoder := json.NewEncoder(buf)
	err := encoder.Encode(fields)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	url := strings.Join([]string{APIURL, "repos", owner, repo, "issues", number}, "/")
	req, err := http.NewRequest("PATCH", url, buf)
	req.SetBasicAuth(os.Getenv("GITHUB_USER"), os.Getenv("GITHUB_PASS"))
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to edit issue: %s", resp.Status)
	}
	var issue github.Issue
	if err = json.NewDecoder(resp.Body).Decode(&issue); err != nil {
		return nil, err
	}
	return &issue, nil
}
