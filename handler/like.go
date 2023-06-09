package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
)

func GetLikes(c echo.Context) error {
	likes := []Like{}

	if err := DB.Find(&likes).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "likes not found")
		}
		return err
	}

	response := map[string]interface{}{
		"likes": likes,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateLike(c echo.Context) error {
	like := Like{}

	if err := c.Bind(&like); err != nil {
		return err
	}

	DB.Create(&like)

	response := map[string]interface{}{
		"like": like,
	}

	return c.JSON(http.StatusCreated, response)
}

func DeleteLike(c echo.Context) error {
	like := new(Like)

	if err := c.Bind(&like); err != nil {
		return err
	}

	if err := DB.Where("user_id = ? AND record_id = ?", like.UserID, like.RecordID).Delete(&like).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	response := map[string]interface{}{
		"likes": like,
	}

	return c.JSON(http.StatusNoContent, response)
}
