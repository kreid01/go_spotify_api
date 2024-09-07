package models

import (
	"gorm.io/gorm"
  	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func ConnectDatabase() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database.")
	}

	err = db.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Playlist{})
	if err != nil {
		panic(err)
	}

       DB = db
}

