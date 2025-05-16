package database

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	dbName := os.Getenv("DB_NAME")
	DB, err = gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
}

func Migrate() {
	err := DB.AutoMigrate(&User{})

	if err != nil {
		panic("failed to migrate database")
	}
}
