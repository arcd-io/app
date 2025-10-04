package http

import (
	"fmt"
	"io"

	"github.com/labstack/echo/v4"
)

func createGithubGroup(e *echo.Group) {
	e.GET("/github/repo/callback", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})

	e.POST("/github/webhook", func(c echo.Context) error {
		// print all headers
		for k, v := range c.Request().Header {
			fmt.Println(k, v)
		}

		buf, err := io.ReadAll(c.Request().Body)
		if err != nil {
			return err
		}

		fmt.Println(string(buf))

		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
