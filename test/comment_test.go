package test

import (
	"bytes"
	"encoding/json"
	"strconv"

	// "errors"
	// "github.com/labstack/echo/v4"
	// "gorm.io/gorm"
	"net/http"
	"net/http/httptest"

	// "strconv"
	"strings"
	"testing"

	// . "bike_noritai_api/handler"
	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
	. "bike_noritai_api/router"
)

func TestGetUserComments(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/users/1/comments", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"record_id":1,"user_name":"tester1","text":"AAAAAAAAAAAAAAA"`

	expectedBody2 := `"id":2,"user_id":1,"record_id":1,"user_name":"tester1","text":"BBBBBBBBBBBBBBBBB"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestGetRecordComments(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/records/2/comments", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":3,"user_id":2,"record_id":2,"user_name":"tester1","text":"CCCCCCCCCCCCCCCCC"`

	expectedBody2 := `"id":4,"user_id":2,"record_id":2,"user_name":"tester1","text":"DDDDDDDDDDDDDDDDDD"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestCreateComment(t *testing.T) {
	router := NewRouter()

	comment := Comment{
		UserName: "Tester",
		Text:     "EEEEEEEEEEEEEEEEEEEE",
	}

	reqBody, err := json.Marshal(comment)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	var userID int64 = 1
	var recordID int64 = 1

	req := httptest.NewRequest(http.MethodPost, "/api/users/"+strconv.Itoa(int(userID))+"/records/"+strconv.Itoa(int(recordID))+"/comments", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody Comment
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resBody.ID)
	}
	if resBody.UserID != userID {
		t.Errorf("expected comment user id to be %v but got %v", userID, resBody.UserID)
	}
	if resBody.UserID != recordID {
		t.Errorf("expected comment record id to be %v but got %v", recordID, resBody.UserID)
	}
	if resBody.UserName != comment.UserName {
		t.Errorf("expected comment user name to be %v but got %v", comment.UserName, resBody.UserName)
	}
	if resBody.Text != comment.Text {
		t.Errorf("expected comment text to be %v but got %v", comment.Text, resBody.Text)
	}
}

func TestUpdateComment(t *testing.T) {
	comment := Comment{
		UserID:   1,
		RecordID: 1,
		UserName: "Tester",
		Text:     "FFFFFFFFFFF",
	}
	if err := DB.Create(&comment).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	updatedComment := Comment{
		ID:       comment.ID,
		UserID:   1,
		RecordID: 1,
		UserName: "Update Tester",
		Text:     "GGGGGGGGGGGG",
	}
	reqBody, err := json.Marshal(updatedComment)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	router := NewRouter()
	req := httptest.NewRequest(http.MethodPatch, "/api/users/"+strconv.Itoa(int(comment.UserID))+"/records/"+strconv.Itoa(int(comment.RecordID))+"/comments/"+strconv.Itoa(int(comment.ID)), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody Comment
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.UserID != updatedComment.UserID {
		t.Errorf("expected comment user id to be %v but got %v", updatedComment.UserID, resBody.UserID)
	}
	if resBody.UserID != updatedComment.RecordID {
		t.Errorf("expected comment record id to be %v but got %v", updatedComment.RecordID, resBody.UserID)
	}
	if resBody.UserName != updatedComment.UserName {
		t.Errorf("expected comment user name to be %v but got %v", updatedComment.UserName, resBody.UserName)
	}
	if resBody.Text != updatedComment.Text {
		t.Errorf("expected comment text to be %v but got %v", updatedComment.Text, resBody.Text)
	}
}

// func TestDeleteSpot(t *testing.T) {
// 	router := NewRouter()
// 	req := httptest.NewRequest(http.MethodDelete, "/api/users/1/spots/1", nil)
// 	res := httptest.NewRecorder()
// 	router.ServeHTTP(res, req)

// 	if res.Code != http.StatusNoContent {
// 		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
// 	}

// 	var deletedSpot Spot
// 	err := DB.First(&deletedSpot, "1").Error
// 	if !errors.Is(err, gorm.ErrRecordNotFound) {
// 		t.Errorf("expected spot record to be deleted, but found: %v", deletedSpot)
// 	}
// }
