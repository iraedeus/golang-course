package usecase

import (
	"golang-course/internal/collector/adapter"
	"golang-course/internal/collector/domain"
)

type CollectorUseCase struct {
	githubAdapter *adapter.GitHubAdapter
}

func NewCollectorUseCase(a *adapter.GitHubAdapter) *CollectorUseCase {
	return &CollectorUseCase{
		githubAdapter: a,
	}
}

func (u *CollectorUseCase) Execute(owner, repoName string) (domain.Repo, error) {
	return u.githubAdapter.GetRepoInfo(owner, repoName)
}
