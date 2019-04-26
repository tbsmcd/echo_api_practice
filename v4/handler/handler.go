package handler

import (
	"net/http"

	"../model"
	"github.com/labstack/echo"
)

// User は JSON を制御する構造体
type User struct {
	Name string `json:"name" xml:"name" form:"name" query:"name"`
}

// GetUser はユーザを返すメソッド http localhost:1323/v4/users/1
func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		id := c.Param("id")
		user, err := model.FindUser(id)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusOK, user)
	}
}

// SaveUser はユーザを保存するメソッド http localhost:1323/v4/users name=tbsmcd
func SaveUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		u := new(User)
		if err := c.Bind(u); err != nil {
			return err
		}
		name := u.Name
		err := model.AddUser(name)
		if err != nil {
			return err
		}
		return c.JSON(http.StatusCreated, u)
	}
}
