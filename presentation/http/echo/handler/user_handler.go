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
func NewUserHandler(u usecase.UserUseCase) UserHandler {
	return &userHandler{u}
}

func Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World! GetUsers")
}
