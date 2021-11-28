package router

import (
  "github.com/a2-ito/go-echo-onion-sample/presentation/http/handler"
)

func SetRouter(e *echo.Echo, h handler.AppHandler) {
	e.GET("/", h.Get)
	e.GET("/users", h.GetUsers)
}
