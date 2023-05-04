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
	// . "bike_noritai_api/repository"
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

// func TestUpdateSpot(t *testing.T) {
// 	spot := Spot{
// 		UserID:      1,
// 		Name:        "東京スカイツリー",
// 		Image:       "http://test.com",
// 		Type:        "観光",
// 		Address:     "〒131-0045 東京都墨田区押上１丁目１−２",
// 		HpURL:       "https://www.tokyo-skytree.jp/",
// 		OpenTime:    "10:00~21:00",
// 		OffDay:      "",
// 		Parking:     true,
// 		Description: "大林建設が施工した日本最高峰の電波塔です。",
// 		Lat:         35.71021159216932,
// 		Lng:         139.81076575474597,
// 	}
// 	if err := DB.Create(&spot).Error; err != nil {
// 		t.Fatalf("failed to create test user: %v", err)
// 	}

// 	updatedSpot := Spot{
// 		ID:          spot.ID,
// 		UserID:      1,
// 		Name:        "豊島美術館",
// 		Image:       "http://test.com",
// 		Type:        "観光",
// 		Address:     "〒761-4662 香川県小豆郡土庄町豊島唐櫃６０７",
// 		HpURL:       "https://benesse-artsite.jp/art/teshima-artmuseum.html",
// 		OpenTime:    "9:00~17:00",
// 		OffDay:      "",
// 		Parking:     true,
// 		Description: "安藤忠雄が設計したユニークな美術館です。",
// 		Lat:         34.49158555200611,
// 		Lng:         134.09277913086976,
// 	}
// 	reqBody, err := json.Marshal(updatedSpot)
// 	if err != nil {
// 		t.Fatalf("failed to marshal request body: %v", err)
// 	}

// 	router := NewRouter()
// 	req := httptest.NewRequest(http.MethodPatch, "/api/users/"+strconv.Itoa(int(spot.UserID))+"/spots/"+strconv.Itoa(int(spot.ID)), bytes.NewBuffer(reqBody))
// 	req.Header.Set("Content-Type", "application/json")
// 	res := httptest.NewRecorder()
// 	router.ServeHTTP(res, req)

// 	if res.Code != http.StatusCreated {
// 		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
// 	}

// 	var resBody Spot
// 	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
// 		t.Fatalf("failed to unmarshal response body: %v", err)
// 	}
// 	if resBody.ID != spot.ID {
// 		t.Errorf("expected spot ID to be %v but got %v", spot.ID, resBody.ID)
// 	}
// 	if resBody.UserID != updatedSpot.UserID {
// 		t.Errorf("expected spot user_id to be %v but got %v", updatedSpot.UserID, resBody.UserID)
// 	}
// 	if resBody.Name != updatedSpot.Name {
// 		t.Errorf("expected spot name to be %v but got %v", updatedSpot.Name, resBody.Name)
// 	}
// 	if resBody.Image != updatedSpot.Image {
// 		t.Errorf("expected spot Image to be %v but got %v", updatedSpot.Image, resBody.Image)
// 	}
// 	if resBody.Type != updatedSpot.Type {
// 		t.Errorf("expected spot Type to be %v but got %v", updatedSpot.Type, resBody.Type)
// 	}
// 	if resBody.Address != updatedSpot.Address {
// 		t.Errorf("expected spot Address to be %v but got %v", updatedSpot.Address, resBody.Address)
// 	}
// 	if resBody.HpURL != updatedSpot.HpURL {
// 		t.Errorf("expected spot HpURL to be %v but got %v", updatedSpot.HpURL, resBody.HpURL)
// 	}
// 	if resBody.OpenTime != updatedSpot.OpenTime {
// 		t.Errorf("expected spot OpenTime to be %v but got %v", updatedSpot.OpenTime, resBody.OpenTime)
// 	}
// 	if resBody.OffDay != updatedSpot.OffDay {
// 		t.Errorf("expected spot OffDay to be %v but got %v", updatedSpot.OffDay, resBody.OffDay)
// 	}
// 	if resBody.Parking != updatedSpot.Parking {
// 		t.Errorf("expected spot Parking to be %v but got %v", updatedSpot.Parking, resBody.Parking)
// 	}
// 	if resBody.Description != updatedSpot.Description {
// 		t.Errorf("expected spot Description to be %v but got %v", updatedSpot.Description, resBody.Description)
// 	}
// 	if resBody.Lat != updatedSpot.Lat {
// 		t.Errorf("expected spot Lat to be %v but got %v", updatedSpot.Lat, resBody.Lat)
// 	}
// 	if resBody.Lng != updatedSpot.Lng {
// 		t.Errorf("expected spot Lng to be %v but got %v", updatedSpot.Lng, resBody.Lng)
// 	}
// }

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
