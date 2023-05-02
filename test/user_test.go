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

func TestGetUsers(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	err := GetUsers(c)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"name":"tester1","email":"tester1@bike_noritai_dev","password":"password","area":"東海","prefecture":"三重県","url":"http://test.com","bike_name":"CBR650R","experience":5`

	expectedBody2 := `"id":2,"name":"tester2","email":"tester2@bike_noritai_dev","password":"password","area":"関東","prefecture":"東京都","url":"http://test.com","bike_name":"CBR1000RR","experience":10`

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

func TestGetUser(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodGet, "/api/users/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v, want %v", res.Code, http.StatusOK)
	}

	expectedBody := `"id":1,"name":"tester1","email":"tester1@bike_noritai_dev","password":"password","area":"東海","prefecture":"三重県","url":"http://test.com","bike_name":"CBR650R","experience":5`

	if !strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}
}

func TestCreateUser(t *testing.T) {
	e := echo.New()

	user := User{
		Name:       "tester3",
		Email:      "tester3@bike_noritai_dev.com",
		Password:   "password",
		Area:       "関西",
		Prefecture: "大阪",
		Url:        "",
		BikeName:   "Ninja650",
		Experience: 10,
	}
	reqBody, err := json.Marshal(user)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	if err := CreateUser(c); err != nil {
		t.Fatalf("failed to create user: %v", err)
	}

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody User
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID == 0 {
		t.Errorf("expected user ID to be non-zero but got %v", resBody.ID)
	}
	if resBody.Name != user.Name {
		t.Errorf("expected user name to be %v but got %v", user.Name, resBody.Name)
	}
	if resBody.Email != user.Email {
		t.Errorf("expected user email to be %v but got %v", user.Email, resBody.Email)
	}
	if resBody.Password != user.Password {
		t.Errorf("expected user password to be %v but got %v", user.Password, resBody.Password)
	}
	if resBody.Area != user.Area {
		t.Errorf("expected user area to be %v but got %v", user.Area, resBody.Area)
	}
	if resBody.Prefecture != user.Prefecture {
		t.Errorf("expected user prefecture to be %v but got %v", user.Prefecture, resBody.Prefecture)
	}
	if resBody.Url != user.Url {
		t.Errorf("expected user url to be %v but got %v", user.Url, resBody.Url)
	}
	if resBody.BikeName != user.BikeName {
		t.Errorf("expected user bike name to be %v but got %v", user.BikeName, resBody.BikeName)
	}
	if resBody.Experience != user.Experience {
		t.Errorf("expected user experience to be %v but got %v", user.Experience, resBody.Experience)
	}
}

func TestUpdateUser(t *testing.T) {
	e := echo.New()

	user := User{
		Name:       "John Doe",
		Email:      "john.doe@example.com",
		Password:   "password",
		Area:       "関西",
		Prefecture: "大阪",
		Url:        "",
		BikeName:   "Ninja650",
		Experience: 10,
	}
	if err := DB.Create(&user).Error; err != nil {
		t.Fatalf("failed to create test user: %v", err)
	}

	updatedUser := User{
		ID:         user.ID,
		Name:       "Jane Smith",
		Email:      "jane.smith@example.com",
		Password:   "update_password",
		Area:       "九州",
		Prefecture: "福岡",
		Url:        "https://example.com",
		BikeName:   "R6",
		Experience: 16,
	}
	reqBody, err := json.Marshal(updatedUser)
	if err != nil {
		t.Fatalf("failed to marshal request body: %v", err)
	}
	req := httptest.NewRequest(http.MethodPut, "/api/users/"+strconv.Itoa(int(user.ID)), bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)
	c.SetParamNames("user_id")
	c.SetParamValues(strconv.Itoa(int(user.ID)))

	if err := UpdateUser(c); err != nil {
		t.Fatalf("failed to update user: %v", err)
	}

	if res.Code != http.StatusCreated {
		t.Errorf("expected status code %v but got %v", http.StatusCreated, res.Code)
	}

	var resBody User
	if err := json.Unmarshal(res.Body.Bytes(), &resBody); err != nil {
		t.Fatalf("failed to unmarshal response body: %v", err)
	}
	if resBody.ID != user.ID {
		t.Errorf("expected user ID to be %v but got %v", user.ID, resBody.ID)
	}
	if resBody.Name != updatedUser.Name {
		t.Errorf("expected user name to be %v but got %v", updatedUser.Name, resBody.Name)
	}
	if resBody.Email != updatedUser.Email {
		t.Errorf("expected user email to be %v but got %v", updatedUser.Email, resBody.Email)
	}
	if resBody.Password != updatedUser.Password {
		t.Errorf("expected user password to be %v but got %v", updatedUser.Password, resBody.Password)
	}
	if resBody.Area != updatedUser.Area {
		t.Errorf("expected user area to be %v but got %v", updatedUser.Area, resBody.Area)
	}
	if resBody.Prefecture != updatedUser.Prefecture {
		t.Errorf("expected user prefecture to be %v but got %v", updatedUser.Prefecture, resBody.Prefecture)
	}
	if resBody.Url != updatedUser.Url {
		t.Errorf("expected user url to be %v but got %v", updatedUser.Url, resBody.Url)
	}
	if resBody.BikeName != updatedUser.BikeName {
		t.Errorf("expected user bike_name to be %v but got %v", updatedUser.BikeName, resBody.BikeName)
	}
	if resBody.Experience != updatedUser.Experience {
		t.Errorf("expected user experience to be %v but got %v", updatedUser.Experience, resBody.Experience)
	}
}

func TestDeleteUser(t *testing.T) {
	router := NewRouter()
	req := httptest.NewRequest(http.MethodDelete, "/api/users/1", nil)
	res := httptest.NewRecorder()
	router.ServeHTTP(res, req)

	if res.Code != http.StatusNoContent {
		t.Errorf("expected status code %v, but got %v", http.StatusNoContent, res.Code)
	}

	var deletedUser User
	err := DB.First(&deletedUser, "1").Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		t.Errorf("expected user record to be deleted, but found: %v", deletedUser)
	}
}
