package main

import (
	//"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	interactor "interactor"
	"net/http"
	router "router"
)

func main() {
	e := echo.New()

	//conn := conf.NewDBConnection()
	i := interactor.NewInteractor()
	h := i.NewAppHandler()

	router.SetRouter(e, h)
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.Logger.Fatal(e.Start(":1323"))
}
