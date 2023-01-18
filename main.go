package main

import (
	"context"

	// "echo-mysql/config"
	"echo-mysql/internal/app"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	app.Route(e, context.Background())
	e.Logger.Fatal(e.Start(":8000"))
}
