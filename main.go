package main

import (
	"context"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// HTTP handlers for callbacks and webhooks
func handleGithubWebhook(c echo.Context) error {
	log.Println("GitHub webhook received")
	// Handle GitHub webhook payload here
	return c.JSON(200, map[string]string{"status": "received"})
}

func handleGithubCallback(c echo.Context) error {
	code := c.QueryParam("code")
	state := c.QueryParam("state")
	log.Printf("GitHub OAuth callback - code: %s, state: %s", code, state)
	// Handle GitHub OAuth callback here
	return c.JSON(200, map[string]string{"status": "authenticated"})
}

func handleHealth(c echo.Context) error {
	return c.JSON(200, map[string]string{"status": "healthy"})
}

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	e.POST("/webhook/github", handleGithubWebhook)
	e.GET("/auth/github/callback", handleGithubCallback)
	e.GET("/health", handleHealth)

	server := &http.Server{
		Addr:    "localhost:8080",
		Handler: h2c.NewHandler(e, &http2.Server{}),
	}

	log.Println("Server starting on localhost:8080")
	log.Fatal(server.ListenAndServe())
}
