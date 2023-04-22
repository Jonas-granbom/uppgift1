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

	DB.Create(&Player{Id: 1, Name: "Jonas Gretzky", TeamId: 5, JerseyNumber: 99, BirthYear: 1985})
	DB.Create(&Player{Id: 3, Name: "Peter Forsberg", TeamId: 6, JerseyNumber: 21, BirthYear: 1973})
	DB.Create(&Player{Id: 4, Name: "Martin Brodeur", TeamId: 7, JerseyNumber: 30, BirthYear: 1985})

	DB.Create(&Team{Id: 5, Name: "Toronto Maple Leafs", City: "Toronto", FoundedYear: 1917})
	DB.Create(&Team{Id: 6, Name: "Colorado Avalanche", City: "Colorado", FoundedYear: 1972})
	DB.Create(&Team{Id: 7, Name: "St. Louis Blues", City: "St. Louis", FoundedYear: 1966})
}
