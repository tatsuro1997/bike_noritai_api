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

	if err := DB.Find(&users).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "users not found")
		}
		return err
	}

	response := map[string]interface{}{
		"users": users,
	}

	return c.JSON(http.StatusOK, response)
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

	response := map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateUser(c echo.Context) error {
	user := User{}

	if err := c.Bind(&user); err != nil {
		return err
	}

	DB.Create(&user)

	response := map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusCreated, response)
}

func UpdateUser(c echo.Context) error {
	user := new(User)

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.First(&user, userID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := DB.Model(&user).Updates(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusCreated, response)
}

func DeleteUser(c echo.Context) error {
	user := new(User)

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("id = ?", userID).Delete(&user).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"user": user,
	}

	return c.JSON(http.StatusNoContent, response)
}
