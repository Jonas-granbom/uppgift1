package data

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	DB, _ = gorm.Open(sqlite.Open("./database/gorm.db"), &gorm.Config{})
	DB.AutoMigrate(&Player{}, &Team{})
}

/* func SeedData() {
	DB.Create(&Player{Id: 1, Name: "Jonas", TeamId: 2, JerseyNumber: 99, BirthYear: 1985})
	DB.Create(&Team{Id: 2, Name: "Toronto Maple Leaves", City: "Toronto", FoundedYear: 1954})
}
 */