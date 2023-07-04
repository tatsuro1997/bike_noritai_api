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
	e.GET("/api/users/:user_id/spots", handler.GetUserSpots)
	e.POST("/api/users/:user_id/spots", handler.CreateSpot)
	e.PATCH("/api/users/:user_id/spots/:spot_id", handler.UpdateSpot)
	e.DELETE("/api/users/:user_id/spots/:spot_id", handler.DeleteSpot)

	e.GET("/api/records", handler.GetRecords)
	e.GET("/api/users/:user_id/records", handler.GetUserRecords)
	e.GET("/api/spots/:spot_id/records", handler.GetSpotRecords)
	e.GET("/api/records/:record_id", handler.GetRecord)
	e.POST("/api/users/:user_id/spots/:spot_id/records", handler.CreateRecord)
	e.PATCH("/api/users/:user_id/spots/:spot_id/records/:record_id", handler.UpdateRecord)
	e.DELETE("/api/users/:user_id/spots/:spot_id/records/:record_id", handler.DeleteRecord)

	e.GET("/api/users/:user_id/comments", handler.GetUserComments)
	e.GET("/api/spots/:spot_id/comments", handler.GetSpotComments)
	e.POST("/api/users/:user_id/records/:record_id/comments", handler.CreateComment)
	e.PATCH("/api/users/:user_id/records/:record_id/comments/:comment_id", handler.UpdateComment)
	e.DELETE("/api/users/:user_id/records/:record_id/comments/:comment_id", handler.DeleteComment)

	// COMFIRMME: unnecessary?
	e.GET("/api/users/:user_id/bookmarks", handler.GetBookmarks)
	e.GET("/api/spots/:spot_id/bookmarks", handler.GetSpotBookmarks)
	e.POST("/api/users/:user_id/spots/:spot_id/bookmarks", handler.CreateBookmark)
	e.DELETE("/api/users/:user_id/spots/:spot_id/bookmarks/:bookmark_id", handler.DeleteBookmark)

	e.GET("/api/likes", handler.GetLikes)
	e.POST("/api/like", handler.CreateLike)
	e.DELETE("/api/like", handler.DeleteLike)

	return e
}
