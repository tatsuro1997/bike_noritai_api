package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
)

func GetSpots(c echo.Context) error {
	spots := []Spot{}

	if err := DB.Find(&spots).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "spots not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, spots)
}
