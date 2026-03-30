package adapter

import (
	"context"

	"golang-course/api/proto"
	"golang-course/gateway/internal/domain"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
)

type CollectorGrpcClient struct {
	grpcClient proto.GithubServiceClient
}

func NewCollectorGrpcClient(address string) (*CollectorGrpcClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &CollectorGrpcClient{
		grpcClient: proto.NewGithubServiceClient(conn),
	}, nil
}

func (c *CollectorGrpcClient) GetRepo(owner, repo string) (domain.Repo, error) {
	req := &proto.RepositoryRequest{
		Owner:    owner,
		RepoName: repo,
	}
	resp, err := c.grpcClient.GetRepository(context.Background(), req)
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return domain.Repo{}, domain.ErrNotFound
		}
		return domain.Repo{}, err
	}

	return domain.Repo{
		Name:        resp.Name,
		Description: resp.Description,
		Stars:       int(resp.Stars),
		Forks:       int(resp.Forks),
		CreatedAt:   resp.CreatedAt,
	}, nil
}
