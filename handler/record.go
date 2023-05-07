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

func GetUserRecords(c echo.Context) error {
	records := []Record{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("user_id = ?", userID).Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "records not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, records)
}

func GetSpotRecords(c echo.Context) error {
	records := []Record{}

	spotID := c.Param("spot_id")
	if spotID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.Where("spot_id = ?", spotID).Find(&records).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "records not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, records)
}

func GetRecord(c echo.Context) error {
	record := []Record{}

	recordID := c.Param("record_id")
	if recordID == "" {
		return c.JSON(http.StatusBadRequest, "record ID is required")
	}

	if err := DB.Where("id = ?", recordID).Find(&record).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "record not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, record)
}

func CreateRecord(c echo.Context) error {
	record := Record{}

	if err := c.Bind(&record); err != nil {
		return err
	}

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	record.UserID = userID
	record.SpotID = spotID

	DB.Create(&record)
	return c.JSON(http.StatusCreated, record)
}

func UpdateRecord(c echo.Context) error {
	record := new(Record)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	recordID := c.Param("record_id")
	if recordID == "" {
		return c.JSON(http.StatusBadRequest, "record ID is required")
	}

	if err := DB.First(&record, recordID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if record.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and record do not match")
	}

	if record.SpotID != spotID {
		return c.JSON(http.StatusBadRequest, "spot and record do not match")
	}

	if err := c.Bind(&record); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := DB.Model(&record).Where("id=?", record.ID).Updates(&record).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, record)
}

func DeleteRecord(c echo.Context) error {
	record := new(Record)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	recordID := c.Param("record_id")
	if recordID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.First(&record, recordID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if record.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and record do not match")
	}

	if record.SpotID != spotID {
		return c.JSON(http.StatusBadRequest, "spot and record do not match")
	}

	if err := DB.Where("id = ?", recordID).Delete(&record).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, record)
}
