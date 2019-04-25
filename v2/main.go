package main

import (
	"./handler"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	g := e.Group("/v2/")
	g.GET("users/:id", handler.GetUser())
	g.POST("users", handler.SaveUser())
	e.Logger.Fatal(e.Start(":1323"))
}
