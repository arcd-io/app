package grpc

import (
	"log"

	"github.com/arcd-io/arcd/gen/github/v1/githubv1connect"
	"github.com/labstack/echo/v4"
)

func NewServer(e *echo.Echo) {
	github := &GithubServer{}
	path, handler := githubv1connect.NewGithubServiceHandler(github)

	e.Any(path+"*", echo.WrapHandler(handler))

	log.Println("Connect-RPC service available at:", path)
}
