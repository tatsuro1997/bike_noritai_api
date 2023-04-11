package main

import (
	rep "bike_noritai_api/repository"
	"bike_noritai_api/router"
)

func main() {
	db, _ := rep.DB.DB()
	defer db.Close()

	e := router.NewRouter()

	e.Logger.Fatal(e.Start(":8080"))
}
