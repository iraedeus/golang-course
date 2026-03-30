package delivery

import (
	"context"
	"errors"

	"golang-course/api/proto"
	"golang-course/collector/internal/domain"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RepoUseCase interface {
	Execute(owner string, repoName string) (domain.Repo, error)
}

type GrpcHandler struct {
	useCase RepoUseCase
	proto.UnimplementedGithubServiceServer
}

func NewGrpcHandler(uc RepoUseCase) *GrpcHandler {
	return &GrpcHandler{
		useCase: uc,
	}
}

func (h *GrpcHandler) GetRepository(ctx context.Context, req *proto.RepositoryRequest) (*proto.RepositoryResponse, error) {
	repo, err := h.useCase.Execute(req.GetOwner(), req.GetRepoName())
	if err != nil {
		if errors.Is(err, domain.ErrNoFound) {
			return nil, status.Error(codes.NotFound, "repository not found")
		}
		return nil, err
	}

	return &proto.RepositoryResponse{
		Name:        repo.Name,
		Description: repo.Description,
		Stars:       int32(repo.Stars),
		Forks:       int32(repo.Forks),
		CreatedAt:   repo.CreatedAt,
	}, nil
}
