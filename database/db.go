package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func NewDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/echo?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	return DB
}

func GetDBInstance() *gorm.DB {
	return DB
}
