package bootstrap

import (
	"github.com/jinzhu/gorm"
	"go-boilerplate/models"
)

func Models(db *gorm.DB) {
	DB := db
	var user models.User
	DB.AutoMigrate(&user)
}
