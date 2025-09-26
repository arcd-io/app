package grpc

import (
	"context"
	"log"

	"connectrpc.com/connect"
	githubv1 "github.com/arcd-io/arcd/gen/github/v1"
)

type GithubServer struct{}

func (r *GithubServer) AddRepository(
	ctx context.Context,
	req *connect.Request[githubv1.AddRepositoryRequest],
) (*connect.Response[githubv1.AddRepositoryResponse], error) {
	log.Println("Request headers: ", req.Header())
	res := connect.NewResponse(&githubv1.AddRepositoryResponse{
		Id: "123",
	})

	res.Header().Set("Github-Version", "v1")
	return res, nil
}
