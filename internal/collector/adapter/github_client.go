package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-course/internal/collector/domain"
)

type repoDTO struct {
	Name        string `json:"full_name"`
	Description string `json:"description"`
	Stars       int    `json:"stargazers_count"`
	Forks       int    `json:"forks_count"`
	CreatedAt   string `json:"created_at"`
}

type GitHubAdapter struct {
	httpClient *http.Client
}

func NewGitHubAdapter() *GitHubAdapter {
	return &GitHubAdapter{
		httpClient: &http.Client{},
	}
}

func (a *GitHubAdapter) GetRepoInfo(owner, repoName string) (domain.Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)

	resp, err := http.Get(url)
	if err != nil {
		return domain.Repo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return domain.Repo{}, fmt.Errorf("repository not found (status code: %d)", resp.StatusCode)
	}

	var repo repoDTO
	err = json.NewDecoder(resp.Body).Decode(&repo)
	if err != nil {
		return domain.Repo{}, err
	}

	return domain.Repo{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       repo.Stars,
		Forks:       repo.Forks,
		CreatedAt:   repo.CreatedAt,
	}, nil
}
