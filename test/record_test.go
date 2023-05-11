package test

import (
	"bytes"
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

func TestGetRecords(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/records", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"spot_id":1,"date":"2023-01-01","weather":"晴れ","temperature":23.4,"running_time":4,"distance":120.4,"description":"最高のツーリング日和でした！"`

	expectedBody2 := `"id":2,"user_id":1,"spot_id":2,"date":"2023-01-01","weather":"曇り","temperature":26.1,"running_time":7,"distance":184.1,"description":"なんとか天気が持って良かったです！"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestGetUserRecords(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/users/1/records", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"spot_id":1,"date":"2023-01-01","weather":"晴れ","temperature":23.4,"running_time":4,"distance":120.4,"description":"最高のツーリング日和でした！"`

	expectedBody2 := `"id":2,"user_id":1,"spot_id":2,"date":"2023-01-01","weather":"曇り","temperature":26.1,"running_time":7,"distance":184.1,"description":"なんとか天気が持って良かったです！"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestGetSpotRecords(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/spots/1/records", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"spot_id":1,"date":"2023-01-01","weather":"晴れ","temperature":23.4,"running_time":4,"distance":120.4,"description":"最高のツーリング日和でした！"`

	expectedBody2 := `"id":3,"user_id":2,"spot_id":1,"date":"2023-01-01","weather":"雨","temperature":13.4,"running_time":2,"distance":50.6,"description":"朝から雨で大変でした。"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}
}

func TestGetRecord(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/records/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"spot_id":1,"date":"2023-01-01","weather":"晴れ","temperature":23.4,"running_time":4,"distance":120.4,"description":"最高のツーリング日和でした！"`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}
}

func TestCreateRecord(t *testing.T) {
	router := NewRouter()

	record := Record{
		Date:        "2023-01-01",
		Weather:     "曇り",
		Temperature: 23.4,
		RunningTime: 4.5,
		Distance:    201.6,
		Description: "AAAAAAAAAAAAAA",
	}

	reqBody, err := json.Marshal(record)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	var userID int64 = 1
	var spotID int64 = 1

	req := httptest.NewRequest(http.MethodPost, "/api/users/"+strconv.Itoa(int(userID))+"/spots/"+strconv.Itoa(int(spotID))+"/records", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody Record
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resBody.ID)
	}
	if resBody.UserID != userID {
		t.Errorf("expected record user id to be %v but got %v", userID, resBody.UserID)
	}
	if resBody.SpotID != spotID {
		t.Errorf("expected record record id to be %v but got %v", spotID, resBody.SpotID)
	}
	if resBody.Date != record.Date {
		t.Errorf("expected record date to be %v but got %v", record.Date, resBody.Date)
	}
	if resBody.Weather != record.Weather {
		t.Errorf("expected record Weather to be %v but got %v", record.Weather, resBody.Weather)
	}
	if resBody.Temperature != record.Temperature {
		t.Errorf("expected record Temperature to be %v but got %v", record.Temperature, resBody.Temperature)
	}
	if resBody.RunningTime != record.RunningTime {
		t.Errorf("expected record RunningTime to be %v but got %v", record.RunningTime, resBody.RunningTime)
	}
	if resBody.Distance != record.Distance {
		t.Errorf("expected record Distance to be %v but got %v", record.Distance, resBody.Distance)
	}
	if resBody.Description != record.Description {
		t.Errorf("expected record Description to be %v but got %v", record.Description, resBody.Description)
	}
}

func TestUpdateRecord(t *testing.T) {
	record := Record{
		UserID:      1,
		SpotID:      1,
		Date:        "2023-01-01",
		Weather:     "曇り",
		Temperature: 23.4,
		RunningTime: 4.5,
		Distance:    201.6,
		Description: "AAAAAAAAAAAAAA",
	}
	if err := DB.Create(&record).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	updatedRecord := Record{
		ID:          record.ID,
		UserID:      1,
		SpotID:      1,
		Date:        "2023-02-01",
		Weather:     "晴れ",
		Temperature: 33.4,
		RunningTime: 6.5,
		Distance:    211.5,
		Description: "BBBBBBBBBBBBB",
	}
	reqBody, err := json.Marshal(updatedRecord)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	router := NewRouter()
	req := httptest.NewRequest(http.MethodPatch, "/api/users/"+strconv.Itoa(int(record.UserID))+"/spots/"+strconv.Itoa(int(record.SpotID))+"/records/"+strconv.Itoa(int(record.ID)), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody Record
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resBody.ID)
	}
	if resBody.UserID != updatedRecord.UserID {
		t.Errorf("expected record user id to be %v but got %v", updatedRecord.UserID, resBody.UserID)
	}
	if resBody.SpotID != updatedRecord.SpotID {
		t.Errorf("expected record record id to be %v but got %v", updatedRecord.SpotID, resBody.SpotID)
	}
	if resBody.Date != updatedRecord.Date {
		t.Errorf("expected record date to be %v but got %v", updatedRecord.Date, resBody.Date)
	}
	if resBody.Weather != updatedRecord.Weather {
		t.Errorf("expected record Weather to be %v but got %v", updatedRecord.Weather, resBody.Weather)
	}
	if resBody.Temperature != updatedRecord.Temperature {
		t.Errorf("expected record Temperature to be %v but got %v", updatedRecord.Temperature, resBody.Temperature)
	}
	if resBody.RunningTime != updatedRecord.RunningTime {
		t.Errorf("expected record RunningTime to be %v but got %v", updatedRecord.RunningTime, resBody.RunningTime)
	}
	if resBody.Distance != updatedRecord.Distance {
		t.Errorf("expected record Distance to be %v but got %v", updatedRecord.Distance, resBody.Distance)
	}
	if resBody.Description != updatedRecord.Description {
		t.Errorf("expected record Description to be %v but got %v", updatedRecord.Description, resBody.Description)
	}
}

func TestDeleteRecord(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodDelete, "/api/users/1/spots/1/records/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
	}

	var deletedRecord *Record
	err := DB.First(&deletedRecord, "1").Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected record to be deleted, but found: %v", deletedRecord)
	}
}
