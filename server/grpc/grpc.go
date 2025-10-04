package grpc

import (
	"log"

	"github.com/arcd-io/arcd/gen/auth/v1/authv1connect"
	"github.com/arcd-io/arcd/gen/github/v1/githubv1connect"
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func NewServer(e *echo.Echo, db *bun.DB) {
	github := &GithubServer{}
	githubPath, githubHandler := githubv1connect.NewGithubServiceHandler(github)
	e.Any(githubPath+"*", echo.WrapHandler(githubHandler))
	log.Println("Connect-RPC service available at:", githubPath)

	auth := NewAuthServer(db)
	authPath, authHandler := authv1connect.NewAuthServiceHandler(auth)
	e.Any(authPath+"*", echo.WrapHandler(authHandler))
	log.Println("Connect-RPC service available at:", authPath)
}
