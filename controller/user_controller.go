package controller

import (
	"goder/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetUser(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	user := model.User{}
	user.FirstById(id)

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	user := model.User{
		Name: name,
	}
	user.Create()

	return c.JSON(http.StatusOK, user)
}
