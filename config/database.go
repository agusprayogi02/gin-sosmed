package config

import (
	"fmt"

	"gin-sosmed/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadDB() {
	// url := strings.Split(ENV.DB_URL, ":")
	dsn := fmt.Sprintf("%v:%v@tcp(%v)/%v?charset=utf8mb4&parseTime=True&loc=Local", ENV.DB_USER, ENV.DB_PASS, ENV.DB_URL, ENV.DB_NAME)
	// dsn := fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v sslmode=disable TimeZone=UTC", url[0], ENV.DB_USER, ENV.DB_PASS, ENV.DB_NAME, url[len(url)-1])
	// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	if err := db.AutoMigrate(&entity.User{}, &entity.Post{}, &entity.Room{}, &entity.Wisma{}, &entity.Customer{}); err != nil {
		panic(err)
	}
	DB = db
}
