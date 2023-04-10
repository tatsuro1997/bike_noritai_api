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

	DB.Migrator().DropTable(&User{})

	DB.AutoMigrate(&User{})
	// DB.AutoMigrate(&Spot{})

	user := []User{
		{
			Name:       "tester1",
			Email:      "tester1@bike_noritai_dev",
			Password:   "password",
			Area:       "東海",
			Prefecture: "三重県",
			Url:        "http://test.com",
			BikeName:   "CBR650R",
			Experience: 5,
		},
		{
			Name:       "tester2",
			Email:      "tester2@bike_noritai_dev",
			Password:   "password",
			Area:       "関東",
			Prefecture: "東京都",
			Url:        "http://test.com",
			BikeName:   "CBR1000RR",
			Experience: 10,
		},
	}
	DB.Create(&user)
}
