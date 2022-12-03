package handler

import (
	"github.com/a2-ito/go-echo-onion-sample/application"
	"github.com/labstack/echo"
	"net/http"
)

type (
	UserHandler interface {
		Get(c echo.Context) error
		GetUsers(c echo.Context) error
	}

	userHandler struct {
		UserUseCase usecase.UserUseCase
	}
)

// func NewUserHandler(u usecase.UserUseCase) UserHandler {
func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func (h *userHandler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (h *userHandler) GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World! GetUsers")
}
