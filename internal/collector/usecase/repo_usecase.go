package usecase

import (
	"golang-course/internal/collector/adapter"
	"golang-course/internal/collector/domain"
)

// CollectorUseCase orchestrates the process of collecting repository information.
// It acts as a bridge between the transport layer (gRPC) and data sources (GitHub).
type CollectorUseCase struct {
	githubAdapter *adapter.GitHubAdapter
}

func NewCollectorUseCase(a *adapter.GitHubAdapter) *CollectorUseCase {
	return &CollectorUseCase{
		githubAdapter: a,
	}
}

// Execute performs the main business logic: requesting data from the adapter
// and returning it to the caller.
func (u *CollectorUseCase) Execute(owner, repoName string) (domain.Repo, error) {
	return u.githubAdapter.GetRepoInfo(owner, repoName)
}
