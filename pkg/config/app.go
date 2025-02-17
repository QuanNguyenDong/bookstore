package config

import (
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

var (
	db * gorm.DB
)

func Connect() {
	dsn := "root:secret@tcp(localhost:3306)/bookstore?charset=utf8&parseTime=True&loc=Local"
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}