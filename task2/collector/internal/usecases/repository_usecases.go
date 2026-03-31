package usecases

import (
	"collector/internal/domain"
	"fmt"
	"regexp"
	"strings"
)

type RepositoryDriver interface {
	GetRepository(authorName, repoName string) (*domain.Repository, error)
}

type RepositoryUsecases struct {
	repositoryDriver RepositoryDriver
}

func NewRepositoryUsecases(repositoryDriver RepositoryDriver) *RepositoryUsecases {
	return &RepositoryUsecases{repositoryDriver}
}

func (ru *RepositoryUsecases) isValidGithubRepoName(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("name is empty")
	}
	if len(name) > 100 {
		return fmt.Errorf("name is too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9-_\.]+$`)
	if !re.MatchString(name) {
		return fmt.Errorf("name contains invalid symbols")
	}

	return nil
}

func (ru *RepositoryUsecases) IsValidGitHubUsername(name string) error {
	if len(name) == 0 {
		return fmt.Errorf("name is empty")
	}
	if len(name) > 39 {
		return fmt.Errorf("name is too long")
	}

	re := regexp.MustCompile(`^[a-zA-Z0-9-]+$`)
	if !re.MatchString(name) {
		return fmt.Errorf("name contains invalid symbols")
	}

	if strings.HasPrefix(name, "-") || strings.HasSuffix(name, "-") {
		return fmt.Errorf("name cannot start or end with '-'")
	}

	if strings.Contains(name, "--") {
		return fmt.Errorf("name cannot contain '--'")
	}

	return nil
}

func (ru *RepositoryUsecases) getRepository(authorName, repoName string) (*domain.Repository, error) {
	err := ru.isValidGithubRepoName(repoName)
	if err != nil {
		return &domain.Repository{}, fmt.Errorf("repository name invalid: %s", err)
	}

	err = ru.IsValidGitHubUsername(authorName)
	if err != nil {
		return &domain.Repository{}, fmt.Errorf("user name invalid: %s", err)
	}

	return ru.repositoryDriver.GetRepository(authorName, repoName)
}
