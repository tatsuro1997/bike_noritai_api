package handler

import (
	"errors"
	"net/http"
	"strconv"

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

func GetSpot(c echo.Context) error {
	spot := Spot{}

	spotID := c.Param("spot_id")
	if spotID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.First(&spot, spotID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "spot not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, spot)
}

func GetUserSpot(c echo.Context) error {
	spots := []Spot{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("user_id = ?", userID).Find(&spots).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "spot not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, spots)
}

func CreateSpot(c echo.Context) error {
	spot := Spot{}

	if err := c.Bind(&spot); err != nil {
		return err
	}

	DB.Create(&spot)
	return c.JSON(http.StatusCreated, spot)
}

func UpdateSpot(c echo.Context) error {
	spot := new(Spot)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID := c.Param("spot_id")
	if spotID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.First(&spot, spotID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if spot.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and spot do not match")
	}

	if err := c.Bind(&spot); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := DB.Model(&spot).Where("id=?", spot.ID).Updates(&spot).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, spot)
}

func DeleteSpot(c echo.Context) error {
	spot := new(Spot)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID := c.Param("spot_id")
	if spotID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.First(&spot, spotID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if spot.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and spot do not match")
	}

	if err := DB.Where("id = ?", spotID).Delete(&spot).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, spot)
}
