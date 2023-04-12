package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
)

func GetUsers(c echo.Context) error {
	users := []User{}
	DB.Find(&users)
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	user := User{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.First(&user, userID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "user not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, user)
}
