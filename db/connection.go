package db

import (
	"log"

	"github.com/glebarez/sqlite" // Pure go SQLite driver, checkout https://github.com/glebarez/sqlite for details
	// "gorm.io/driver/sqlite"      // Sqlite driver based on GGO
	"gorm.io/gorm"
)

var DB *gorm.DB

func DBConnection() {
	// Connection to database
	var err error
	DB, err = gorm.Open(sqlite.Open("./gorm.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Connected to database")
	}
}
