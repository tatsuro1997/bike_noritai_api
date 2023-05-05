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

func GetBookmarks(c echo.Context) error {
	bookmarks := []Bookmark{}

	userID := c.Param("user_id")
	if userID == "" {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	if err := DB.Where("user_id = ?", userID).Find(&bookmarks).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.JSON(http.StatusNotFound, "bookmarks not found")
		}
		return err
	}

	return c.JSON(http.StatusOK, bookmarks)
}

func CreateBookmark(c echo.Context) error {
	bookmark := Bookmark{}

	if err := c.Bind(&bookmark); err != nil {
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

	bookmark.UserID = userID
	bookmark.SpotID = spotID

	DB.Create(&bookmark)
	return c.JSON(http.StatusCreated, bookmark)
}

func DeleteBookmark(c echo.Context) error {
	bookmark := new(Bookmark)

	userID, _ := strconv.ParseInt(c.Param("user_id"), 10, 64)
	if userID == 0 {
		return c.JSON(http.StatusBadRequest, "user ID is required")
	}

	spotID, _ := strconv.ParseInt(c.Param("spot_id"), 10, 64)
	if spotID == 0 {
		return c.JSON(http.StatusBadRequest, "spot ID is required")
	}

	bookmarkID := c.Param("bookmark_id")
	if bookmarkID == "" {
		return c.JSON(http.StatusBadRequest, "bookmark ID is required")
	}

	if err := DB.First(&bookmark, bookmarkID).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	if bookmark.UserID != userID {
		return c.JSON(http.StatusBadRequest, "user and bookmark do not match")
	}

	if bookmark.SpotID != spotID {
		return c.JSON(http.StatusBadRequest, "spot and bookmark do not match")
	}

	if err := DB.Where("id = ?", bookmarkID).Delete(&bookmark).Error; err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusNoContent, bookmark)
}
