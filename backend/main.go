package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Static("/static", "static")

	e.GET("/user", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{"name": "Jhon Doe"})
	})

	e.Start(":8000")
}
