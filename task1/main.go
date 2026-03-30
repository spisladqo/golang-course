package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

const reposUrl = "https://api.github.com/repos/"

type Repository struct {
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	Parent struct {
		FullName string `json:"full_name"`
	} `json:"parent"`
	Name            string `json:"name"`
	Id              int    `json:"id"`
	StarsCount      int    `json:"stargazers_count"`
	ForksCount      int    `json:"forks_count"`
	CreatedAt       string `json:"created_at"`
	IsFork          bool   `json:"fork"`
	OpenIssuesCount int    `json:"open_issues_count"`
}

func printRepoStats(repo Repository) {
	isForkYN := ""
	if repo.IsFork {
		isForkYN = "yes"
	} else {
		isForkYN = "no"
	}
	fmt.Println("=Repository statistics=")
	fmt.Println("Name:        ", repo.Name)
	fmt.Println("Owner login: ", repo.Owner.Login)
	fmt.Println("Stars count: ", repo.StarsCount)
	fmt.Println("Forks count: ", repo.ForksCount)
	fmt.Println("Created at:  ", repo.CreatedAt)
	fmt.Println("Is a fork:   ", isForkYN)
	if repo.IsFork {
		fmt.Println("Forked from: ", repo.Parent.FullName)
	}
	fmt.Println("Open issues: ", repo.OpenIssuesCount)
}

func sendRequest(ownerLogin, repoName string) (*http.Response, error) {
	url := reposUrl + ownerLogin + "/" + repoName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)

	return resp, err
}

func runLoop() {
	for {
		var ownerLogin string
		var repoName string

		fmt.Print("Enter repository owner login: ")
		fmt.Scanln(&ownerLogin)
		fmt.Print("Enter repository name: ")
		fmt.Scanln(&repoName)
		fmt.Println()

		resp, err := sendRequest(ownerLogin, repoName)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println("Try again")
			fmt.Println()
			continue
		} else if resp.StatusCode != http.StatusOK {
			fmt.Println("Http error: ", resp.Status)
			fmt.Println("Try again")
			fmt.Println()
			continue
		}

		var repo Repository
		err = json.NewDecoder(resp.Body).Decode(&repo)
		if err != nil {
			fmt.Println("Error when marshalling:", err)
			continue
		}
		resp.Body.Close()

		printRepoStats(repo)
		fmt.Println()
	}
}

func main() {
	runLoop()
}
