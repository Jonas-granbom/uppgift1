package data

import (
	"os"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	s := os.Getenv("RUNENVIRONMENT")
	var filePath = ""
	if s == "Production" {
		filePath = "/database/gorm.sqlite"
	} else {
		filePath = "database/gorm.sqlite"
	}

	os.MkdirAll("database", 0700)
	DB, _ = gorm.Open(sqlite.Open(filePath), &gorm.Config{})
	DB.AutoMigrate(&Player{}, &Team{})

}

 func SeedData() {
	DB.Create(&Player{Id: 1, Name: "Jonas", TeamId: 2, JerseyNumber: 99, BirthYear: 1985})
	DB.Create(&Team{Id: 2, Name: "Toronto Maple Leaves", City: "Toronto", FoundedYear: 1954})
}

