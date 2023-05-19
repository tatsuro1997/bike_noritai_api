package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"strconv"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	. "bike_noritai_api/handler"
	. "bike_noritai_api/model"
	. "bike_noritai_api/repository"
	. "bike_noritai_api/router"
)

func TestGetSpots(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/spots", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := GetSpots(c)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"name":"豊受大神宮 (伊勢神宮 外宮）","image":"","type":"観光","address":"三重県伊勢市豊川町２７９","hp_url":"https://www.isejingu.or.jp/about/geku/","open_time":"5:00~18:00","off_day":"","parking":true,"description":"外宮から行くのが良いみたいですよ。","lat":34.4879,"lng":136.704`

	expectedBody2 := `"id":2,"user_id":1,"name":"伊勢神宮（内宮）","image":"","type":"観光","address":"三重県伊勢市宇治館町１","hp_url":"https://www.isejingu.or.jp/","open_time":"5:00~18:00","off_day":"","parking":true,"description":"日本最大の由緒正しき神社です。","lat":34.4562,"lng":136.726`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if !strings.Contains(res.Body.String(), expectedBody2) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody2)
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}

func TestGetSpot(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/spots/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"user_id":1,"name":"豊受大神宮 (伊勢神宮 外宮）","image":"","type":"観光","address":"三重県伊勢市豊川町２７９","hp_url":"https://www.isejingu.or.jp/about/geku/","open_time":"5:00~18:00","off_day":"","parking":true,"description":"外宮から行くのが良いみたいですよ。","lat":34.487865,"lng":136.70374`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}
}

func TestGetUserSpots(t *testing.T) {
	spot := Spot{
		UserID:      2,
		Name:        "東京スカイツリー",
		Image:       "http://test.com",
		Type:        "観光",
		Address:     "〒131-0045 東京都墨田区押上１丁目１−２",
		HpURL:       "https://www.tokyo-skytree.jp/",
		OpenTime:    "10:00~21:00",
		OffDay:      "",
		Parking:     true,
		Description: "大林建設が施工した日本最高峰の電波塔です。",
		Lat:         35.71021159216932,
		Lng:         139.81076575474597,
	}
	if err := DB.Create(&spot).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/users/2/spots", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":3,"user_id":2,"name":"東京スカイツリー","image":"http://test.com","type":"観光","address":"〒131-0045 東京都墨田区押上１丁目１−２","hp_url":"https://www.tokyo-skytree.jp/","open_time":"10:00~21:00","off_day":"","parking":true,"description":"大林建設が施工した日本最高峰の電波塔です。","lat":35.710213,"lng":139.81076`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}
}

func TestCreateSpot(t *testing.T) {
	router := NewRouter()

	spot := Spot{
		UserID:      1,
		Name:        "東京スカイツリー",
		Image:       "http://test.com",
		Type:        "観光",
		Address:     "〒131-0045 東京都墨田区押上１丁目１−２",
		HpURL:       "https://www.tokyo-skytree.jp/",
		OpenTime:    "10:00~21:00",
		OffDay:      "",
		Parking:     true,
		Description: "大林建設が施工した日本最高峰の電波塔です。",
		Lat:         35.71021159216932,
		Lng:         139.81076575474597,
	}

	reqBody, err := json.Marshal(spot)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/users/1/spots", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	resBody := ResponseBody{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), &resBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
		return
	}

	resSpot := resBody.Spot
	if resSpot.ID == 0 {
		t.Errorf("expected spot ID to be non-zero but got %v", resSpot.ID)
	}
	if resSpot.UserID != spot.UserID {
		t.Errorf("expected spot user id to be %v but got %v", spot.UserID, resSpot.UserID)
	}
	if resSpot.Name != spot.Name {
		t.Errorf("expected spot name to be %v but got %v", spot.Name, resSpot.Name)
	}
	if resSpot.Image != spot.Image {
		t.Errorf("expected spot image to be %v but got %v", spot.Image, resSpot.Image)
	}
	if resSpot.Type != spot.Type {
		t.Errorf("expected spot type to be %v but got %v", spot.Type, resSpot.Type)
	}
	if resSpot.Address != spot.Address {
		t.Errorf("expected spot address to be %v but got %v", spot.Address, resSpot.Address)
	}
	if resSpot.HpURL != spot.HpURL {
		t.Errorf("expected spot HP URL to be %v but got %v", spot.HpURL, resSpot.HpURL)
	}
	if resSpot.OpenTime != spot.OpenTime {
		t.Errorf("expected spot open time to be %v but got %v", spot.OpenTime, resSpot.OpenTime)
	}
	if resSpot.OffDay != spot.OffDay {
		t.Errorf("expected spot off day to be %v but got %v", spot.OffDay, resSpot.OffDay)
	}
	if resSpot.Parking != spot.Parking {
		t.Errorf("expected spot parking to be %v but got %v", spot.Parking, resSpot.Parking)
	}
	if resSpot.Description != spot.Description {
		t.Errorf("expected spot description to be %v but got %v", spot.Description, resSpot.Description)
	}
	if resSpot.Lat != spot.Lat {
		t.Errorf("expected spot lat to be %v but got %v", spot.Lat, resSpot.Lat)
	}
	if resSpot.Lng != spot.Lng {
		t.Errorf("expected spot lng to be %v but got %v", spot.Lng, resSpot.Lng)
	}
}

func TestUpdateSpot(t *testing.T) {
	spot := Spot{
		UserID:      1,
		Name:        "東京スカイツリー",
		Image:       "http://test.com",
		Type:        "観光",
		Address:     "〒131-0045 東京都墨田区押上１丁目１−２",
		HpURL:       "https://www.tokyo-skytree.jp/",
		OpenTime:    "10:00~21:00",
		OffDay:      "",
		Parking:     true,
		Description: "大林建設が施工した日本最高峰の電波塔です。",
		Lat:         35.71021159216932,
		Lng:         139.81076575474597,
	}
	if err := DB.Create(&spot).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	updatedSpot := Spot{
		ID:          spot.ID,
		UserID:      1,
		Name:        "豊島美術館",
		Image:       "http://test.com",
		Type:        "観光",
		Address:     "〒761-4662 香川県小豆郡土庄町豊島唐櫃６０７",
		HpURL:       "https://benesse-artsite.jp/art/teshima-artmuseum.html",
		OpenTime:    "9:00~17:00",
		OffDay:      "",
		Parking:     true,
		Description: "安藤忠雄が設計したユニークな美術館です。",
		Lat:         34.49158555200611,
		Lng:         134.09277913086976,
	}
	reqBody, err := json.Marshal(updatedSpot)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}

	router := NewRouter()
	req := httptest.NewRequest(http.MethodPatch, "/api/users/"+strconv.Itoa(int(spot.UserID))+"/spots/"+strconv.Itoa(int(spot.ID)), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	resBody := ResponseBody{}
	if err := json.Unmarshal([]byte(res.Body.Bytes()), &resBody); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
		return
	}

	resSpot := resBody.Spot

	// var resBody Spot
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resSpot.ID != spot.ID {
		t.Errorf("expected spot ID to be %v but got %v", spot.ID, resSpot.ID)
	}
	if resSpot.UserID != updatedSpot.UserID {
		t.Errorf("expected spot user_id to be %v but got %v", updatedSpot.UserID, resSpot.UserID)
	}
	if resSpot.Name != updatedSpot.Name {
		t.Errorf("expected spot name to be %v but got %v", updatedSpot.Name, resSpot.Name)
	}
	if resSpot.Image != updatedSpot.Image {
		t.Errorf("expected spot Image to be %v but got %v", updatedSpot.Image, resSpot.Image)
	}
	if resSpot.Type != updatedSpot.Type {
		t.Errorf("expected spot Type to be %v but got %v", updatedSpot.Type, resSpot.Type)
	}
	if resSpot.Address != updatedSpot.Address {
		t.Errorf("expected spot Address to be %v but got %v", updatedSpot.Address, resSpot.Address)
	}
	if resSpot.HpURL != updatedSpot.HpURL {
		t.Errorf("expected spot HpURL to be %v but got %v", updatedSpot.HpURL, resSpot.HpURL)
	}
	if resSpot.OpenTime != updatedSpot.OpenTime {
		t.Errorf("expected spot OpenTime to be %v but got %v", updatedSpot.OpenTime, resSpot.OpenTime)
	}
	if resSpot.OffDay != updatedSpot.OffDay {
		t.Errorf("expected spot OffDay to be %v but got %v", updatedSpot.OffDay, resSpot.OffDay)
	}
	if resSpot.Parking != updatedSpot.Parking {
		t.Errorf("expected spot Parking to be %v but got %v", updatedSpot.Parking, resSpot.Parking)
	}
	if resSpot.Description != updatedSpot.Description {
		t.Errorf("expected spot Description to be %v but got %v", updatedSpot.Description, resSpot.Description)
	}
	if resSpot.Lat != updatedSpot.Lat {
		t.Errorf("expected spot Lat to be %v but got %v", updatedSpot.Lat, resSpot.Lat)
	}
	if resSpot.Lng != updatedSpot.Lng {
		t.Errorf("expected spot Lng to be %v but got %v", updatedSpot.Lng, resSpot.Lng)
	}
}

func TestDeleteSpot(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodDelete, "/api/users/1/spots/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
	}

	var deletedSpot Spot
	err := DB.First(&deletedSpot, "1").Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected spot record to be deleted, but found: %v", deletedSpot)
	}
}
