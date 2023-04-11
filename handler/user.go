package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
)

func GetUsers(c echo.Context) error {
	users := []User{}
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}
