package models

import "gorm.io/gorm"

var DB *gorm.DB

func Init(db *gorm.DB) {
	DB = db
	var user User
	DB.AutoMigrate(&user)
}
