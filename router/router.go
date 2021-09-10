package router

import (
	"github.com/labstack/echo/v4"
	"goder/controller"
	"net/http"
)

func Init() *echo.Echo {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	/*
	e.GET("/api/product/:id", controller.GetProduct)
	e.POST("/api/product", controller.CreateProduct)
	e.PUT("/api/product/:id", controller.UpdateProduct)
	e.DELETE("/api/product/:id", controller.DeleteProduct)

	 */

	e.GET("/api/user", controller.GetUserALL)
	e.POST("/api/user", controller.CreateUser)
	e.GET("/api/user/:id", controller.GetUserByID)
	e.PATCH("/api/user/:id", controller.UpdateUser)

	e.GET("/api/order/:id", controller.GetOrderByID)
	e.POST("api/order", controller.CreateOrder)
	e.PATCH("api/order/:id", controller.UpdateOrder)
	return e
}