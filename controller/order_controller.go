package controller

import (
	"goder/model"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetOrder(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	order := model.Order{}
	order.FirstById(id)

	return c.JSON(http.StatusOK, order)
}

func CreateOrder(c echo.Context) error {
	uid, _ := strconv.Atoi(c.FormValue("user_id"))
	userID := uint(uid)
	user := model.User{}
	user.FirstById(userID)

	order := model.Order{
		User: user,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	}
	order.Create()

	return c.JSON(http.StatusOK, order)
}
