package test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo/v4"

	. "bike_noritai_api/handler"
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

	expectedBody := `[{"id":1,"name":"tester1","email":"tester1@bike_noritai_dev","password":"password","area":"東海","prefecture":"三重県","url":"http://test.com","bike_name":"CBR650R","experience":5,"posts":null},{"id":2,"name":"tester2","email":"tester2@bike_noritai_dev","password":"password","area":"関東","prefecture":"東京都","url":"http://test.com","bike_name":"CBR1000RR","experience":10,"posts":null}]`

	if strings.Contains(res.Body.String(), expectedBody) {
		t.Errorf("unexpected response body: got %v, want %v", res.Body.String(), expectedBody)
	}

	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
}
