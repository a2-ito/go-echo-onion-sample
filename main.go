package main

import (
	"github.com/a2-ito/go-echo-rest-samples/presentation/http/echo/router"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()
  i := interactor.NewInteractor()
  router.SetRouter(e, )
  /*
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":1323"))
  */
}
