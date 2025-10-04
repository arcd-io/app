package http

import (
	"github.com/labstack/echo/v4"
	"github.com/uptrace/bun"
)

func createAuthGroup(e *echo.Group, db *bun.DB) {
	authGroup := e.Group("/auth")

	authGroup.GET("/github", func(c echo.Context) error {
	})

	authGroup.GET("/github/callback", func(c echo.Context) error {
	})
}
