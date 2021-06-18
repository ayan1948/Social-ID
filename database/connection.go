package database

import (
	"github.com/mymodules/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Connect() {
	connection, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	connection.AutoMigrate(&models.User{})
}
