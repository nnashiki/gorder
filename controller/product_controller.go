package controller

import (
	"goder/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func GetProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	product := model.Product{}
	product.FirstById(id)

	return c.JSON(http.StatusOK, product)
}

func CreateProduct(c echo.Context) error {
	name := c.FormValue("name")
	p, _ := strconv.Atoi(c.FormValue("price"))
	price := uint(p)

	product := model.Product{
		Name:  name,
		Price: price,
	}
	product.Create()

	return c.JSON(http.StatusOK, product)
}

func UpdateProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	name := c.FormValue("name")
	p, _ := strconv.Atoi(c.FormValue("price"))
	price := uint(p)

	product := model.Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
	product.Updates()

	return c.JSON(http.StatusOK, product)
}

func DeleteProduct(c echo.Context) error {
	i, _ := strconv.Atoi(c.Param("id"))
	id := uint(i)
	product := model.Product{}
	product.DeleteById(id)

	return c.JSON(http.StatusOK, product)
}