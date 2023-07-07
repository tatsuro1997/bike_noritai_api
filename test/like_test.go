package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gorm.io/gorm"

	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
	. "bike_noritai_api/router"
)

func TestGetLikes(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/likes", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"record_id":1`

	expectedBody2 := `"id":2,"user_id":1,"record_id":2`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestCreateLike(t *testing.T) {
	router := NewRouter()

	like := Like{
		UserID:   3,
		RecordID: 3,
	}

	reqBody, err := json.Marshal(like)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/like", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	resBody := ResponseLikeBody{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}

	resLike := resBody.Like
	if resLike.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resLike.ID)
	}
	if resLike.UserID != like.UserID {
		t.Errorf("expected comment user id to be %v but got %v", like.UserID, resLike.UserID)
	}
	if resLike.RecordID != like.RecordID {
		t.Errorf("expected comment record id to be %v but got %v", like.RecordID, resLike.RecordID)
	}
}

func TestDeleteLike(t *testing.T) {
	router := NewRouter()

	like := Like{
		UserID:   1,
		RecordID: 1,
	}

	reqBody, err := json.Marshal(&like)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodDelete, "/api/like", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
	}

	var deletedLike *Like
	err = DB.First(&deletedLike, like.ID).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected like record to be deleted, but found: %v", deletedLike)
	}
}
