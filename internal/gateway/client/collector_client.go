package client

import (
	"context"

	"golang-course/api/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type CollectorClient struct {
	grpcClient proto.GithubServiceClient
}

func NewCollectorClient(address string) (*CollectorClient, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return &CollectorClient{
		grpcClient: proto.NewGithubServiceClient(conn),
	}, nil
}

func (c *CollectorClient) GetRepo(owner, repo string) (*proto.RepositoryResponse, error) {
	req := &proto.RepositoryRequest{
		Owner:    owner,
		RepoName: repo,
	}
	resp, err := c.grpcClient.GetRepository(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
