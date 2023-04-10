package repository

import (
	"log"

	. "bike_noritai_api/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func init() {
	dsn := "tester:password@tcp(db:3306)/bike_noritai_dev?charset=utf8mb4&parseTime=True&loc=Local"

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(dsn + "database can't connect")
	}

	DB.AutoMigrate(&User{})
	// DB.AutoMigrate(&Spot{})
}
