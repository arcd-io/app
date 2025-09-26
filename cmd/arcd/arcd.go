package main

import (
	"github.com/arcd-io/arcd/server/grpc"
	"github.com/arcd-io/arcd/server/http"
)

func main() {
	// database := database.New()
	// githubService := github.NewService(database)

	e := http.NewServer()
	grpc.NewServer(e)
}
