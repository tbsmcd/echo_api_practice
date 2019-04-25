package main

import (
	"./db"
	"./handler"

	"github.com/labstack/echo"
)

func main() {
	db.Init()
	e := echo.New()
	g := e.Group("/v3/")
	g.GET("users/:id", handler.GetUser())
	g.POST("users", handler.SaveUser())
	e.Logger.Fatal(e.Start(":1323"))
}
