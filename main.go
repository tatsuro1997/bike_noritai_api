package main

import (
	. "bike_noritai_api/repository"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.CORS())

	db, _ := DB.DB()
	defer db.Close()

	e.Logger.Fatal(e.Start(":8080"))
}
