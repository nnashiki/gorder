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

	e.GET("/api/product/:id", controller.GetProduct)
	e.POST("/api/product", controller.CreateProduct)
	e.PUT("/api/product/:id", controller.UpdateProduct)
	e.DELETE("/api/product/:id", controller.DeleteProduct)

	e.GET("/api/user/:id", controller.GetUser)
	e.POST("/api/user", controller.CreateUser)
	e.PATCH("/api/user/:id", controller.UpdateUser)

	e.GET("/api/order/:id", controller.GetOrder)
	e.POST("api/order", controller.CreateOrder)
	return e
}