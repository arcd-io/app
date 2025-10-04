package grpc

import (
	"context"

	"connectrpc.com/connect"
	authv1 "github.com/arcd-io/arcd/gen/auth/v1"
	"github.com/uptrace/bun"
)

type AuthServer struct {
	db *bun.DB
}

func NewAuthServer(db *bun.DB) *AuthServer {
	return &AuthServer{db: db}
}

func (s *AuthServer) GetSession(
	ctx context.Context,
	req *connect.Request[authv1.GetSessionRequest],
) (*connect.Response[authv1.GetSessionResponse], error) {
	return nil, nil
}
