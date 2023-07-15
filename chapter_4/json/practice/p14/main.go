package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Issue struct {
	Title       string `json:"title"`
	State       string `json:"state"`
	Description string `json:"body"`
	Milestone   struct {
		Title string `json:"title"`
	} `json:"milestone"`
	User struct {
		Login string `json:"login"`
	} `json:"user"`
}

func main() {
	http.HandleFunc("/issues", handleIssues)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
func handleIssues(w http.ResponseWriter, r *http.Request) {
	owner := r.URL.Query().Get("owner")
	repo := r.URL.Query().Get("repo")

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/issues", owner, repo)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	var issues []Issue
	err = json.NewDecoder(resp.Body).Decode(&issues)
	if err != nil {
		log.Fatal(err)
	}
	for _, issue := range issues {
		fmt.Fprintf(w, "Title: %s\n", issue.Title)
		fmt.Fprintf(w, "State: %s\n", issue.State)
		fmt.Fprintf(w, "Description: %s\n", issue.Description)
		fmt.Fprintf(w, "Milestone: %s\n", issue.Milestone.Title)
		fmt.Fprintf(w, "User: %s\n\n", issue.User.Login)
		fmt.Fprintln(w, "==========================")
	}
}
