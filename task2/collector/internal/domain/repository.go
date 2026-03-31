package domain

import "time"

type Repository struct {
	Owner           Owner     `json:"owner"`
	Parent          Parent    `json:"parent"`
	Name            string    `json:"name"`
	Id              int       `json:"id"`
	StarsCount      int       `json:"stargazers_count"`
	ForksCount      int       `json:"forks_count"`
	CreatedAt       time.Time `json:"created_at"`
	IsFork          bool      `json:"fork"`
	OpenIssuesCount int       `json:"open_issues_count"`
}

type Owner struct {
	Login string `json:"login"`
}

type Parent struct {
	FullName string `json:"full_name"`
}
