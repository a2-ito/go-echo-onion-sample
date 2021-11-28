package handler

import (
	"github.com/a2-ito/go-echo-onion-sample/usecase"
	"github.com/labstack/echo"
	"net/http"
)

type (
	UserHandler interface {
		Get(c echo.Context) error
		GetUsers(c echo.Context) error
	}
)

//func NewUserHandler(u usecase.UserUseCase) UserHandler {
func NewUserHandler() UserHandler {
	return &userHandler
}

func (h *userHandler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (h *userHandler) GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World! GetUsers")
}
