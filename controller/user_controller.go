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

func UpdateUser(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	email := c.FormValue("email")

	user := model.User{}
	user.FirstById(id)
	userNew := model.User{Email: email}
	user.Updates(userNew)

	return c.JSON(http.StatusOK, user)
}