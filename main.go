package main

import (
	"github.com/a2-ito/go-echo-onion-sample/interactor"
	"github.com/a2-ito/go-echo-onion-sample/presentation/http/echo/router"
	//"github.com/nanamen/go-echo-rest-sample/interactor"
	"github.com/labstack/echo"
	"net/http"
)

func main() {
	e := echo.New()

	//i := interactor.NewInteractor()
	//h := i.NewAppHandler()

	//router.SetRouter(e, h)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
