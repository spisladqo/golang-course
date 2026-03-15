package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Repository struct {
	Owner struct {
		Login string `json:"login"`
	} `json:"owner"`
	Name       string `json:"name"`
	Id         int    `json:"id"`
	StarsCount int    `json:"stargazers_count"`
	ForksCount int    `json:"forks_count"`
}

func main() {
	url := "https://api.github.com/repos/spisladqo/golang-course"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	client := &http.Client{Timeout: 30 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Printf("unexpected status: got %v\n", resp.Status)
		return
	}

	var repo Repository
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("repo name:   ", repo.Name)
	fmt.Println("repo id:     ", repo.Id)
	fmt.Println("owner login: ", repo.Owner.Login)
	fmt.Println("stars count: ", repo.StarsCount)
	fmt.Println("forks count: ", repo.ForksCount)
}
