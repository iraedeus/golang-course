package handler

import (
	"context"

	"golang-course/api/proto"
	"golang-course/internal/collector/usecase"
)

type GrpcHandler struct {
	collectorUC *usecase.CollectorUseCase
	proto.UnimplementedGithubServiceServer
}

func NewGrpcHandler(uc *usecase.CollectorUseCase) *GrpcHandler {
	return &GrpcHandler{
		collectorUC: uc,
	}
}

// GetRepository - это реализация метода из .proto файла
func (h *GrpcHandler) GetRepository(ctx context.Context, req *proto.RepositoryRequest) (*proto.RepositoryResponse, error) {
	repo, err := h.collectorUC.Execute(req.GetOwner(), req.GetRepoName())
	if err != nil {
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
