package models

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var DB *gorm.DB

func ConnectDatabase() {
	host := "viaduct.proxy.rlwy.net"
	user := "postgres"
	password := "d-c42b5d-ABea*4a-fBBCDA*23E1AAe-"
	dbname := "railway"
	port := "45094"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require TimeZone=Asia/Jakarta", host, user, password, dbname, port)
	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Order{})
	database.AutoMigrate(&Item{})

	DB = database
}