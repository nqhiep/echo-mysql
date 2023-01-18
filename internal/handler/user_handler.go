package handler

import (
	"echo-mysql/internal/model"
	"echo-mysql/internal/service"

	"github.com/labstack/echo/v4"
	"net/http"
)

type UserHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) All(c echo.Context) error {
	users, err := h.service.All(c.Request().Context())
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, users)
}

func (h *UserHandler) Load(c echo.Context) error {
	id := c.Param("id")
	user, err := h.service.Load(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())

	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Insert(c echo.Context) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	res, err1 := h.service.Insert(c.Request().Context(), &user)
	if err1 != nil {
		return c.String(http.StatusInternalServerError, err1.Error())

	}
	return c.JSON(http.StatusCreated, res)
}

func (h *UserHandler) Update(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return err
	}

	err := h.service.Update(c.Request().Context(), &user)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}

func (h *UserHandler) Delete(c echo.Context) error {
	id := c.Param("id")
	err := h.service.Delete(c.Request().Context(), id)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, "Deleted!")
}
