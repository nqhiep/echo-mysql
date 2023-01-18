package app

import (
	"context"

	"github.com/labstack/echo/v4"
)

func Route(e *echo.Echo, ctx context.Context) error {
	app, err := NewApp(ctx)
	if err != nil {
		return err
	}

	e.GET("/users", app.UserHandler.All)
	e.GET("/users/:id", app.UserHandler.Load)
	e.POST("users", app.UserHandler.Insert)
	e.PUT("users/:id", app.UserHandler.Update)
	e.DELETE("users/:id", app.UserHandler.Delete)

	return nil
}
