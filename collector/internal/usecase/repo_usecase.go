package usecase

import (
	"golang-course/collector/internal/domain"
)

type RepoProvider interface {
	GetRepoInfo(owner string, repoName string) (domain.Repo, error)
}

type CollectorUseCase struct {
	provider RepoProvider
}

func NewCollectorUseCase(p RepoProvider) *CollectorUseCase {
	return &CollectorUseCase{
		provider: p,
	}
}

func (u *CollectorUseCase) Execute(owner, repoName string) (domain.Repo, error) {
	return u.provider.GetRepoInfo(owner, repoName)
}
