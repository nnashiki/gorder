package controller

import (
	"context"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"goder/models"
	"goder/mysql"
	"net/http"
	"strconv"
)

func GetUserByID(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	ctx := context.Background()
	user, err := models.Users(
		qm.Where("id=?", id),
	).One(ctx, mysql.DB)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func GetUserALL(c echo.Context) error {
	name := c.Param("name")

	ctx := context.Background()
	user, err := models.Users(
		qm.Where("name=?", name),
	).All(ctx, mysql.DB)

	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}


func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	ctx := context.Background()
	user := models.User{
		Name:  name,
		Email: email,
	}
	err := user.Insert(ctx, mysql.DB, boil.Infer())
	if err != nil {
		fmt.Println(err)
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusNotFound,err)
	}
	id := uint(i)
	email := c.FormValue("email")
	ctx := context.Background()
	user, err := models.Users(
		qm.Where("id=?", id),
	).One(ctx, mysql.DB)
	user.Email = email
	user.Update(ctx, mysql.DB, boil.Infer())

	return c.JSON(http.StatusOK, user)
}
