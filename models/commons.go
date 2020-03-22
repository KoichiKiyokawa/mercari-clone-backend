package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Connect() *gorm.DB {
	db, err := gorm.Open("sqlite3", "data.db")
	if err != nil {
		panic("Failed to connect database")
	}
	db.AutoMigrate(&Item{})
	return db
}
