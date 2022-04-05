package database

import (
	"go_admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	database, err := gorm.Open(mysql.Open("root:5sntVLe69LSkrMK@tcp(db:3306)/go_admin?parseTime=true"), &gorm.Config{})

	if err != nil {
		panic("Could not connect to database")
	}

	DB = database

	database.AutoMigrate(&models.User{}, &models.Product{}, &models.Job{}, &models.Message{}, &models.Productimg{}, &models.Comment{})
}
