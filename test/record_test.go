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

	resBody := ResponseRecordBody{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), &resBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
		return
	}

	resRecord := resBody.Record
	if resRecord.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resRecord.ID)
	}
	if resRecord.UserID != userID {
		t.Errorf("expected record user id to be %v but got %v", userID, resRecord.UserID)
	}
	if resRecord.SpotID != spotID {
		t.Errorf("expected record record id to be %v but got %v", spotID, resRecord.SpotID)
	}
	if resRecord.Date != record.Date {
		t.Errorf("expected record date to be %v but got %v", record.Date, resRecord.Date)
	}
	if resRecord.Weather != record.Weather {
		t.Errorf("expected record Weather to be %v but got %v", record.Weather, resRecord.Weather)
	}
	if resRecord.Temperature != record.Temperature {
		t.Errorf("expected record Temperature to be %v but got %v", record.Temperature, resRecord.Temperature)
	}
	if resRecord.RunningTime != record.RunningTime {
		t.Errorf("expected record RunningTime to be %v but got %v", record.RunningTime, resRecord.RunningTime)
	}
	if resRecord.Distance != record.Distance {
		t.Errorf("expected record Distance to be %v but got %v", record.Distance, resRecord.Distance)
	}
	if resRecord.Description != record.Description {
		t.Errorf("expected record Description to be %v but got %v", record.Description, resRecord.Description)
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

	resBody := ResponseRecordBody{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), &resBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
		return
	}

	resRecord := resBody.Record
	if resRecord.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resRecord.ID)
	}
	if resRecord.UserID != updatedRecord.UserID {
		t.Errorf("expected record user id to be %v but got %v", updatedRecord.UserID, resRecord.UserID)
	}
	if resRecord.SpotID != updatedRecord.SpotID {
		t.Errorf("expected record record id to be %v but got %v", updatedRecord.SpotID, resRecord.SpotID)
	}
	if resRecord.Date != updatedRecord.Date {
		t.Errorf("expected record date to be %v but got %v", updatedRecord.Date, resRecord.Date)
	}
	if resRecord.Weather != updatedRecord.Weather {
		t.Errorf("expected record Weather to be %v but got %v", updatedRecord.Weather, resRecord.Weather)
	}
	if resRecord.Temperature != updatedRecord.Temperature {
		t.Errorf("expected record Temperature to be %v but got %v", updatedRecord.Temperature, resRecord.Temperature)
	}
	if resRecord.RunningTime != updatedRecord.RunningTime {
		t.Errorf("expected record RunningTime to be %v but got %v", updatedRecord.RunningTime, resRecord.RunningTime)
	}
	if resRecord.Distance != updatedRecord.Distance {
		t.Errorf("expected record Distance to be %v but got %v", updatedRecord.Distance, resRecord.Distance)
	}
	if resRecord.Description != updatedRecord.Description {
		t.Errorf("expected record Description to be %v but got %v", updatedRecord.Description, resRecord.Description)
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
