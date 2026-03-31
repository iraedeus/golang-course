package usecase

import (
	"time"

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

func formatData(date string) string {
	t, err := time.Parse(time.RFC3339, date)
	if err != nil {
		return date
	}
	return t.Format("15:04:05 02.01.06")
}

func (u *CollectorUseCase) Execute(owner, repoName string) (domain.Repo, error) {
	repo, err := u.provider.GetRepoInfo(owner, repoName)
	if err != nil {
		return repo, err
	}

	repo.CreatedAt = formatData(repo.CreatedAt)
	return repo, err
}
