package handler

import (
	"context"

	"golang-course/api/proto"
	"golang-course/internal/collector/usecase"
)

// GrpcHandler implements the gRPC service defined in the proto file.
// It handles incoming gRPC requests and delegates work to the UseCase.
type GrpcHandler struct {
	collectorUC *usecase.CollectorUseCase
	proto.UnimplementedGithubServiceServer
}

func NewGrpcHandler(uc *usecase.CollectorUseCase) *GrpcHandler {
	return &GrpcHandler{
		collectorUC: uc,
	}
}

// GetRepository is the gRPC method that receives a request,
// calls the business logic, and returns a gRPC-compatible response.
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
