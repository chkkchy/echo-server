package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/hoge", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	//e.Logger.Fatal(e.Start(":1323"))
	e.Logger.Fatal(e.Start(":8080"))
}
