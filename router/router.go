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
	e.DELETE("/api/users/:user_id", handler.DeleteUser)

	e.GET("/api/spots", handler.GetSpots)
	e.GET("/api/spots/:spot_id", handler.GetSpot)
	e.GET("/api/users/:user_id/spots", handler.GetUserSpot)
	e.POST("/api/spots", handler.CreateSpot)
	e.PATCH("/api/users/:user_id/spots/:spot_id", handler.UpdateSpot)
	e.DELETE("/api/users/:user_id/spots/:spot_id", handler.DeleteSpot)

	e.GET("/api/users/:user_id/comments", handler.GetUserComments)
	e.GET("/api/records/:record_id/comments", handler.GetRecordComments)
	e.POST("/api/users/:user_id/records/:record_id/comments", handler.CreateComment)

	return e
}
