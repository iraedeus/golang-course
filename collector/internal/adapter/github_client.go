package adapter

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang-course/collector/internal/domain"
)

type GitHubAdapter struct {
	httpClient *http.Client
}

func NewGitHubAdapter() *GitHubAdapter {
	return &GitHubAdapter{
		httpClient: &http.Client{},
	}
}

func (a *GitHubAdapter) GetRepoInfo(owner string, repoName string) (domain.Repo, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/%s", owner, repoName)

	resp, err := a.httpClient.Get(url)
	if err != nil {
		return domain.Repo{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		if resp.StatusCode == http.StatusNotFound {
			return domain.Repo{}, domain.ErrNotFound
		}
		return domain.Repo{}, fmt.Errorf("github api error: %d", resp.StatusCode)
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
