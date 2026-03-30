package usecase

import "golang-course/gateway/internal/domain"

type CollectorProvider interface {
	GetRepo(owner string, repoName string) (domain.Repo, error)
}

type GatewayUseCase struct {
	provider CollectorProvider
}

func NewGatewayUseCase(p CollectorProvider) *GatewayUseCase {
	return &GatewayUseCase{
		provider: p,
	}
}

func (u *GatewayUseCase) Execute(owner string, repoName string) (domain.Repo, error) {
	return u.provider.GetRepo(owner, repoName)
}
