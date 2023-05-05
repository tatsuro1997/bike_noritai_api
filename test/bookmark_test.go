package test

import (
	"encoding/json"
	"errors"
	"gorm.io/gorm"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
	. "bike_noritai_api/router"
)

func TestGetBookmarks(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/users/1/bookmarks", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"spot_id":1`
	expectedBody2 := `"id":2,"user_id":1,"spot_id":2,`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestCreateBookmark(t *testing.T) {
	router := NewRouter()

	var userID int64 = 3
	var spotID int64 = 3

	req := httptest.NewRequest(http.MethodPost, "/api/users/"+strconv.Itoa(int(userID))+"/spots/"+strconv.Itoa(int(spotID))+"/bookmarks", nil)
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody Bookmark
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resBody.ID)
	}
	if resBody.UserID != userID {
		t.Errorf("expected comment user id to be %v but got %v", userID, resBody.UserID)
	}
	if resBody.SpotID != spotID {
		t.Errorf("expected comment record id to be %v but got %v", spotID, resBody.UserID)
	}
}

func TestDeleteBookmark(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodDelete, "/api/users/1/spots/1/bookmarks/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
	}

	var deletedBookmark *Bookmark
	err := DB.First(&deletedBookmark, "1").Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected spot record to be deleted, but found: %v", deletedBookmark)
	}
}
