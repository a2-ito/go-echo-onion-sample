package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

type UserHander interface {
	Get(c echo.Context) error
	GetUsers(c echo.Context) error
}

func (h *userHnadler) Get(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World!")
}

func (h *userHnadler) GetUsers(c echo.Context) error {
	return c.String(http.StatusOK, "Hello World! GetUsers")
}
