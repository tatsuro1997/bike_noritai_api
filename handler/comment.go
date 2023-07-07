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

func GetUserComments(c echo.Context) error {
	comments := []Comment{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("user_id = ?", userID).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "comments not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, comments)
}

func GetSpotComments(c echo.Context) error {
	comments := []Comment{}

	spotID := c.Param("spot_id")
	if spotID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.Where("spot_id = ?", spotID).Find(&comments).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "comments not found")
		}
		return err
	}

	response := map[string]interface{}{
		"comments": comments,
	}

	return c.JSON(http.StatusOK, response)
}

func CreateComment(c echo.Context) error {
	comment := Comment{}

	if err := c.Bind(&comment); err != nil {
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

	comment.UserID = userID
	comment.SpotID = spotID

	DB.Create(&comment)
	return c.JSON(http.StatusCreated, comment)
}

func UpdateComment(c echo.Context) error {
	comment := new(Comment)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	commentID := c.Param("comment_id")
	if commentID == "" {
		return c.JSON(http.StatusBadRequest, "comment ID is required")
	}

	if err := DB.First(&comment, commentID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if comment.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and comment do not match")
	}

	if comment.SpotID != spotID {
		return c.JSON(http.StatusBadRequest, "spot and comment do not match")
	}

	if err := c.Bind(&comment); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if err := DB.Model(&comment).Where("id=?", comment.ID).Updates(&comment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusCreated, comment)
}

func DeleteComment(c echo.Context) error {
	comment := new(Comment)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	commentID := c.Param("comment_id")
	if commentID == "" {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	if err := DB.First(&comment, commentID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if comment.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and comment do not match")
	}

	if comment.SpotID != spotID {
		return c.JSON(http.StatusBadRequest, "spot and comment do not match")
	}

	if err := DB.Where("id = ?", commentID).Delete(&comment).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, comment)
}
