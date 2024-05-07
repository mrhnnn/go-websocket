package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/cek", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{"Tai": "Kucing"})
	})

	e.Logger.Fatal(e.Start(":8000"))
}
