package main

import (
	"net/http"

	"github.com/labstack/echo"
)

// User は JSON を制御する構造体
type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
}

func main() {
	e := echo.New()
	e.GET("/v1/users/:id", getUser)
	e.POST("/v1/users", saveUser)
	e.Logger.Fatal(e.Start(":1323"))
}

// http localhost:1323/v1/users/1
func getUser(c echo.Context) error {
	id := c.Param("id")
	user := map[string]string{
		"id":   id,
		"name": "tbsmcd",
	}
	return c.JSON(http.StatusOK, user)
}

// http localhost:1323/v1/users name=tbsmcd
func saveUser(c echo.Context) error {
	u := new(User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
}
