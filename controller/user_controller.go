package controller

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	_ "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"goder/models"
	"goder/mysql"
	"net/http"
	"strconv"
	"fmt"
)

func GetUser(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	ctx := context.Background()
	db := mysql.New()
	user, err := models.Users(
		qm.Where("id=?",id),
		).One(ctx, db)

	if err != nil {
		fmt.Println(err)
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func CreateUser(c echo.Context) error {
	name := c.FormValue("name")
	ctx := context.Background()
	db := mysql.New()
	user := models.User{
		Name: name,
		Email: "hoge",
	}
	err := user.Insert(ctx, db, boil.Infer())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, user)
}

func UpdateUser(c echo.Context) error {
	i, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return err
	}
	id := uint(i)
	print(id)
	email := c.FormValue("email")
	ctx := context.Background()
	db := mysql.New()
	user, err := models.Users().One(ctx, db)
	user.Email = email
	user.Update(ctx, db, boil.Infer())

	return c.JSON(http.StatusOK, user)
}
