package router

import (
	"bike_noritai_api/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func NewRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())

	e.GET("/api/users", handler.GetUsers)
	e.GET("/api/users/:user_id", handler.GetUser)
	e.POST("/api/users", handler.CreateUser)
	e.PATCH("/api/users/:user_id", handler.UpdateUser)

	return e
}
