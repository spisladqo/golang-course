package drivers

import (
	"collector/internal/domain"
	"encoding/json"
	"fmt"
	"net/http"
)

type RepositoryDriver struct {
	client *http.Client
}

func NewRepositoryDriver(client *http.Client) *RepositoryDriver {
	return &RepositoryDriver{client}
}

func (rd *RepositoryDriver) sendRequest(req *http.Request) (*http.Response, error) {
	return rd.client.Do(req)
}

func (rd *RepositoryDriver) newGetRepositoryRequest(authorLogin, repoName string) (*http.Request, error) {
	url := "https://api.github.com/repos/" + authorLogin + "/" + repoName

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Request{}, err
	}

	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2026-03-10")

	return req, nil
}

func (rd *RepositoryDriver) GetRepository(authorLogin, repoName string) (*domain.Repository, error) {
	req, err := rd.newGetRepositoryRequest(authorLogin, repoName)
	if err != nil {
		return &domain.Repository{}, err
	}

	resp, err := rd.sendRequest(req)
	if err != nil {
		return &domain.Repository{}, err
	} else if resp.StatusCode != http.StatusOK {
		return &domain.Repository{}, fmt.Errorf("http status is not ok: %d", resp.StatusCode)
	}

	defer func() {
		resp.Body.Close()
	}()

	var repo *domain.Repository
	err = json.NewDecoder(resp.Body).Decode(repo)
	if err != nil {
		return &domain.Repository{}, err
	}

	return repo, nil
}
