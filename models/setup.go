package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := "viaduct.proxy.rlwy.net"
	user := "postgres"
	password := "d-c42b5d-ABea*4a-fBBCDA*23E1AAe-"
	dbname := "railway"
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=5432 sslmode=require TimeZone=Asia/Jakarta"
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Order{})
	database.AutoMigrate(&Item{})

	DB = database
}