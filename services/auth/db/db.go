package db

import (
	"auth/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Init() {
	dsn := "host=localhost user=admin password=admin dbname=pract8 port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal(err)
	}

	DB = db
}
