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
	DB.Migrator().DropTable(&Spot{})

	DB.AutoMigrate(&User{})
	DB.AutoMigrate(&Spot{})

	users := []User{
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
	DB.Create(&users)

	spots := []Spot{
		{
			UserID:      1,
			Name:        "豊受大神宮 (伊勢神宮 外宮）",
			Image:       "http://test.com",
			Type:        "観光",
			Address:     "三重県伊勢市豊川町２７９",
			HpURL:       "https://www.isejingu.or.jp/about/geku/",
			OpenTime:    "5:00~18:00",
			OffDay:      "",
			Parking:     true,
			Description: "外宮から行くのが良いみたいですよ。",
			Lat:         34.48786428571363,
			Lng:         136.70372958477844,
		},
		{
			UserID:      1,
			Name:        "伊勢神宮（内宮）",
			Image:       "http://test.com",
			Type:        "観光",
			Address:     "三重県伊勢市宇治館町１",
			HpURL:       "https://www.isejingu.or.jp/",
			OpenTime:    "5:00~18:00",
			OffDay:      "",
			Parking:     true,
			Description: "日本最大の由緒正しき神社です。",
			Lat:         34.45616423029016,
			Lng:         136.7258739014393,
		},
	}
	DB.Create(&spots)
}
